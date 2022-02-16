package main

import (
	"github.com/mad-development/go-ff/bar"
	"github.com/mad-development/go-ff/foo"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {

	if IsEnabled("endpoint-ab-test") {
		foo.GetFoo(w, r)
		return
	}

	bar.GetBar(w, r)
	return
}
