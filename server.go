package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 8888, "port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		remoteAddr := r.RemoteAddr
		fmt.Fprintf(w, "Hello, your ip: %s\n", remoteAddr)
		forwords := r.Header.Get("X-Forwarded-For")
		fmt.Fprintf(w, "Hello, your forwords: %s\n", forwords)
		cookie := r.Cookies()
		fmt.Fprintf(w, "Hello, your cookie: %s\n", cookie)
	})

	portStr := fmt.Sprintf(":%d", *portPtr)
	log.Println("start listen", portStr)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
