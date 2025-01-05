package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

type app struct {
	addr string
}

func (a app) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func (a app) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All our users")
}

func (a app) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user created")
}

func main() {
	// ANSI escape codes for colors
	const (
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Reset  = "\033[0m"
	)

	_, path, _, ok := runtime.Caller(0)
	if ok {
		log.Printf(Red+"starting `%s`...\n"+Reset, filepath.Base(path))
		app := app{addr: ":8080"}
		mux := http.NewServeMux()

		mux.HandleFunc("GET /", app.indexHandler)
		mux.HandleFunc("GET /users", app.getUsersHandler)
		mux.HandleFunc("POST /users", app.createUsersHandler)
		srv := &http.Server{
			Addr:    app.addr,
			Handler: mux,
		}
		log.Fatal(srv.ListenAndServe())
	}
}
