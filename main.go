package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

// ANSI escape codes for colors
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func main() {

	_, path, _, ok := runtime.Caller(0)
	if ok {
		log.Printf(Red+"starting `%s`...\n"+Reset, filepath.Base(path))
		api := &api{addr: ":8080"}
		mux := http.NewServeMux()

		mux.HandleFunc("GET /users", api.getUsersHandler)
		mux.HandleFunc("POST /users", api.addUserHandler)
		srv := &http.Server{
			Addr:    api.addr,
			Handler: mux,
		}
		log.Fatal(srv.ListenAndServe())
	}
}
