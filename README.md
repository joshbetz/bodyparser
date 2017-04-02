# bodyparser

Automatically parse the net/http request.Body into request.Form data depending on the Content-Type header.

## Example

```go
func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := fmt.Sprintf("Hello %s!", r.FormValue("name"))
		w.Write([]byte(body))
	})

	handler := bodyparser.Middleware(h)
	http.ListenAndServe(":3000", handler)
}
```

## API

```go
func Middleware(h http.Handler) http.Handler
```

A middleware for net/http to parse request bodies into request.Form data
depending on the content type.

```go
func Parse(r *http.Request) (map[string]interface{}, error)
```

Similar to `Middleware`, but instead of fitting the data into request.Form,
returns a `map[string]interface{}`. Useful for parsing more complex data than
string to string.

```go
func ParseJSON(r *http.Request) (map[string]interface{}, error)
```

The underlying JSON parser. Useful if you only care about parsing
application/json data.
