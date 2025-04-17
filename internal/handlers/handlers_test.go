package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name              string
	url               string
	method            string
	params            []postData
	expectedStatuCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2023-10-01"},
		{key: "end", value: "2023-10-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2023-10-01"},
		{key: "end", value: "2023-10-02"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Prateek"},
		{key: "last_name", value: "Singh"},
		{key: "email", value: "p@gmail.com"},
		{key: "phone", value: "1234567890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatuCode {
				t.Errorf("for %s, expected %d, got %d", e.name, e.expectedStatuCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				fmt.Printf("Error in post form: %v", e.name)
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatuCode {
				t.Errorf("for %s, expected %d, got %d", e.name, e.expectedStatuCode, resp.StatusCode)
			}
		}
	}
}
