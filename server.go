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

		if cookie, err := r.Cookie("Flavor"); err != nil {
			http.SetCookie(w, &http.Cookie{Name: "Flavor", Value: "Chocolate Chip"})
		} else {
			cookie.Value = "Oatmeal Raisin"
			http.SetCookie(w, cookie)
		}
		cookies := r.Cookies()
		fmt.Fprintf(w, "Hello, your cookie: %s\n", cookies)

		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		remoteAddr := r.RemoteAddr
		fmt.Fprintf(w, "Hello, your ip: %s\n", remoteAddr)

		forwords := r.Header.Get("X-Forwarded-For")
		fmt.Fprintf(w, "Hello, your forwords: %s\n", forwords)

		ua := r.Header.Get("User-Agent")
		fmt.Fprintf(w, "Hello, your ua: %s\n", ua)
	})

	portStr := fmt.Sprintf(":%d", *portPtr)
	log.Println("start listen", portStr)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
