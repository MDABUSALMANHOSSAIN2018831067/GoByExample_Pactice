package main

import (
	"fmt"
	"net/http"
	"time"
)

var indexPage = `
<!DOCTYPE html>
<html>
<body style="text-align:center">
<h1> My First Heading </h1>
<p> My First Paragraph </p>
</body>
</html>
`

var userInfo = `{
	"name":"salman",
	"age":21

}`

func main() {
	port := ":5000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexPage))
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userInfo))
	})

	s := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Start server  ", port)

	panic(s.ListenAndServe())
}
