# Limit - limit io.ReadCloser [![Go Report Card](https://goreportcard.com/badge/github.com/fharding1/limit)](https://goreportcard.com/report/github.com/fharding1/limit)

## Who

[Franklin Harding](https://harding.coffee/).

## What

A package that provides a way to `limit` the amount of bytes read from an
`io.ReadCloser`.

## Why

The `io` package only provides a `io.LimitedReader`, it doesn't provide a `LimitedReadCloser`, which can be very useful for working with http, files,
etc, etc.

## Example

Say you want to create a middleware that limits any handler from reading
over 10 bytes. You remember the awesome `io.LimitReader` function, but then
realize `Body` is a `ReadCloser`, not just a `Reader`! With `limit` we can
easily implement this middleware, because it has `LimitedReadCloser`, which
has an underlying `io.ReadCloser`. For a more complete example see [here](https://gist.github.com/fharding1/92eb4a65492836933d2972f7b5897633).

```golang
func limitMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                r.Body = limit.ReadCloser(r.Body, 10)
                next.ServeHTTP(w, r)
        })
}
```
