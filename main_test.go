package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

const expected = "hello world"

func TestDefaultViewHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:3333", nil)
	w := httptest.NewRecorder()

	defaultView(w, req)
	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Fatal("wrong response body:", body)
	}
}
