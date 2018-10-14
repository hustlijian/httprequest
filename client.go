package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlPtr := flag.String("url", "http://127.0.0.1:8888", "url to request")
	flag.Parse()

	client := &http.Client{}
	req, err := http.NewRequest("GET", *urlPtr, nil)

	// set cookie
	cookie := http.Cookie{Name: "cookie_name", Value: "lijian"}
	req.AddCookie(&cookie)

	// set ip
	req.Header.Add("X-Forwarded-For", `10.0.0.1`)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data) // recieve data
}
