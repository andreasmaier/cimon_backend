package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
)

func JobUpdates(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)
}