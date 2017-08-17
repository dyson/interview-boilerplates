package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRouteHello(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	base, _ := url.Parse(srv.URL)
	path, _ := url.Parse("/hello")
	res, err := http.Get(base.ResolveReference(path).String())
	if err != nil {
		t.Fatal("could not send GET request:", err)
	}

	expectedStatus := http.StatusOK
	if res.StatusCode != expectedStatus {
		t.Errorf("expected status: %d, got: %s", expectedStatus, res.Status)
	}

	expectedType := "application/json"
	if res.Header.Get("content-type") != expectedType {
		t.Errorf("expected content-type: %s, got: %s", expectedType, res.Header.Get("content-type"))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("could not read response body:", err)
	}

	expectedBody, _ := json.Marshal(map[string]string{"hello": "world"})
	if string(body) != string(expectedBody) {
		t.Errorf("expected body: %s, got: %s", expectedBody, body)
	}
}
