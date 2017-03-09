Got - HTTP API Calls
====================

Package got is a super simple net/http wrapper that does the right thing
for most JSON REST APIs specifically adding:

 * Retry logic with exponential back-off,
 * Reading of body with a MaxSize to avoid running out of memory,
 * Timeout after 30 seconds.

```go
g := got.New()
response, err := g.NewRequest("PUT", url, []byte("...")).Send()
if err == nil {
  // handle error
}

// Make a GET request
response, err := g.NewRequest("GET", url, nil).Send()
// or short hand
response, err := g.Get(url).Send()

// Send JSON in a request
request := g.NewRequest("PUT", url, []byte(`{"key": "value"}`))
request.Header.Set("Content-Type": "application/json")
// or short hand
request := g.Put(url, nil)
err := request.JSON(map[string]{"key": "value"})

// Disable retries for a request
request := g.NewRequest("POST", url, []byte(`{"key": "value"}`))
request.Retries = 0  // Defaults to 5

// Disable retries for all requests
g.Retries = 0
```

For more details see: [godoc.org/github.com/taskcluster/go-got](https://godoc.org/github.com/taskcluster/go-got)
