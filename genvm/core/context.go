package core

import (
	"bytes"
	"fmt"

	"github.com/spacemeshos/go-scale"
)

// Context servers 2 purposes:
// - maintains changes to the system state, that will be applied only after succeful execution
// - accumulates set of reusable objects and data
type Context struct {
	Loader   AccountLoader
	Handler  Handler
	Template Template

	Account   Account
	Principal Address
	Method    uint8

	Header Header
	Args   scale.Encodable

	// consumed and transfered is for MaxGas/MaxSpend validation
	consumed   uint64
	transfered uint64

	order   []Address
	changed map[Address]*Account
}

// Spawn account.
// TODO(dshulyak) only self-spawn is supported for now.
func (c *Context) Spawn(template Address, args scale.Encodable) error {
	principal := ComputePrincipal(template, c.Header.Nonce, args)
	if principal != c.Principal {
		return ErrSpawn
	}

	c.Account.Template = &template
	return nil
}

// Transfer amount to the address after validation passes.
func (c *Context) Transfer(to Address, amount uint64) error {
	if amount > c.Account.Balance {
		return ErrNoBalance
	}
	// noop. only gas is consumed
	if c.Account.Address == to {
		return nil
	}
	c.transfered += amount
	if c.transfered > c.Header.MaxSpend {
		return fmt.Errorf("%w: %d", ErrMaxSpend, c.Header.MaxSpend)
	}

	if c.changed == nil {
		c.changed = map[Address]*Account{}
	}
	account, exist := c.changed[to]
	if !exist {
		loaded, err := c.Loader.Get(to)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInternal, err.Error())
		}
		c.order = append(c.order, to)
		c.changed[to] = &loaded
		account = &loaded
	}

	c.Account.Balance -= amount
	account.Balance += amount
	return nil
}

// Consume gas from the account after validation passes.
func (c *Context) Consume(gas uint64) error {
	amount := gas * c.Header.GasPrice
	if amount > c.Account.Balance {
		return ErrNoBalance
	}
	c.consumed += amount
	if c.consumed > c.Header.MaxGas {
		return ErrMaxGas
	}
	c.Account.Balance -= amount
	return nil
}

// Apply modifified state to the account updater.
func (c *Context) Apply(updater AccountUpdater) error {
	buf := bytes.NewBuffer(nil)
	encoder := scale.NewEncoder(buf)
	c.Template.EncodeScale(encoder)

	c.Account.Nonce = c.Header.Nonce.Counter
	c.Account.State = buf.Bytes()
	if err := updater.Update(c.Account); err != nil {
		return fmt.Errorf("%w: %s", ErrInternal, err.Error())
	}
	for _, address := range c.order {
		account := c.changed[address]
		if err := updater.Update(*account); err != nil {
			return fmt.Errorf("%w: %s", ErrInternal, err.Error())
		}
	}
	return nil
}