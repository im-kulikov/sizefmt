sizefmt
=======

Fork and fixed from [code.cloudfoundry.org/bytefmt](http://code.cloudfoundry.org/bytefmt)

Human-readable byte formatter.

Example:

```go
sizefmt.ByteSize(100.5*sizefmt.MEGABYTE) // returns "100.5M"
sizefmt.ByteSize(uint64(1024)) // returns "1K"
```

For documentation, please see http://godoc.org/github.com/im-kulikov/sizefmt
