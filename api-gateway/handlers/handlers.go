package handlers

import (
	"io/ioutil"
	"net/http"
)

var ListBusinesses = func(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:7777/api/businesses")

	if err != nil {
		// do something
	}

	body, err := ioutil.ReadAll(resp.Body)

	w.Write([]byte(body))

}
