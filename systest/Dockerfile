FROM golang:1.19-alpine as build
RUN apk add libc6-compat gcc musl-dev
WORKDIR /build/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go test -failfast -v -c -o /build/tests.test ./systest/tests/

FROM alpine
COPY --from=build /build/tests.test /bin/tests
