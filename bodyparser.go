package bodyparser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := Parse(r)

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err)
			return
		}

		if r.PostForm == nil || len(r.PostForm) <= 0 {
			if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
				r.PostForm, err = parsePostJson(res)
			} else {
				r.PostForm = make(url.Values)
			}
		}

		// Copy PostForm to Form and parse the query string
		if r.Form == nil || len(r.Form) <= 0 {
			r.Form = nil
			r.ParseForm()
		}

		h.ServeHTTP(w, r)
	})
}

func Parse(r *http.Request) (map[string]interface{}, error) {
	res := map[string]interface{}{}

	// application/x-www-form-urlencoded
	err := r.ParseForm()

	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	if len(r.Form) > 0 {
		for k, v := range r.Form {
			res[k] = interface{}(v)
		}

		return res, nil
	}

	// application/json
	res, err = ParseJSON(r)

	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func ParseJSON(r *http.Request) (map[string]interface{}, error) {
	if r.Header["Content-Type"][0] != "application/json" {
		return nil, nil
	}

	if len(r.Header["Content-Length"]) > 0 && r.Header["Content-Length"][0] == "0" {
		return nil, nil
	}

	reader := io.LimitReader(r.Body, int64(10<<20)) // 10MB
	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func parsePostJson(data map[string]interface{}) (url.Values, error) {
	values := make(url.Values)

	for k, v := range data {
		switch t := v.(type) {
		case string:
			values.Set(k, t)
		case float64:
			values.Set(k, strconv.FormatFloat(t, 'f', -1, 64))
		case bool:
			if t {
				values.Set(k, "true")
			} else {
				values.Set(k, "false")
			}
		case interface{}:
			j, err := json.Marshal(v)

			if err != nil {
				return nil, err
			}

			values.Set(k, string(j))
		}
	}

	return values, nil
}
