package controller

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping())
	mux.HandleFunc("/", Crud())
	return mux
}
