package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	urlPtr := flag.String("url", "http://127.0.0.1:8888", "url to request")
	flag.Parse()

	u, err := url.Parse(*urlPtr)
	if err != nil {
		log.Fatal(err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	// set cookie
	cookie := http.Cookie{Name: "cookie_name", Value: "lijian san"}
	//req.AddCookie(&cookie)

	jar.SetCookies(u, []*http.Cookie{&cookie})

	client := &http.Client{
		Jar: jar,
	}
	req, err := http.NewRequest("GET", u.String(), nil)

	// set ip
	req.Header.Add("X-Forwarded-For", `10.0.0.1`)
	// set ua
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")

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

	fmt.Println("After request cookies:")
	for _, cookie := range jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
}
