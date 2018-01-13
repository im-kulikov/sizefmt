sizefmt
=======

Fork and fixed from [code.cloudfoundry.org/bytefmt](http://code.cloudfoundry.org/bytefmt)

Human-readable byte formatter.

Example:

```go
sizefmt.ByteSize(100.5*sizefmt.MEGABYTE) // returns "100.5M"
sizefmt.ByteSize(uint64(1024)) // returns "1K"
```

```go
sizefmt.Time(time.Now()) // now
sizefmt.Time(time.Now().Add(+2 * time.Hour)) // 2 hours from now
sizefmt.RelTime(time.Unix(0, 0), time.Unix(7*24*60*60, -1), "ago", "") // 6 days ago
```

For documentation, please see http://godoc.org/github.com/im-kulikov/sizefmt
