package test

import (
	"net/http"
	"net/http/httptest"
)

func MockServer(htpStatus int, respBody []byte) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if htpStatus != 0 {
			w.WriteHeader(htpStatus)
			_, _ = w.Write(respBody)
		}
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}
