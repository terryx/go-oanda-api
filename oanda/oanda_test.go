package oanda

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func response(filename string) http.HandlerFunc {
	stub := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		body, _ := ioutil.ReadFile(filename)
		_, err := w.Write(body)
		if err != nil {
			log.Fatal(err)
		}
	})

	return stub
}

func StubResponse(filename string) (api *Api) {
	stub := httptest.NewServer(response(filename))
	api = &Api{
		"oanda-api-key",
		"oanda-account-id",
		stub.URL,
	}

	return api
}
