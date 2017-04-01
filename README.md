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
