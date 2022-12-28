package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/realIP", irlIP)
	http.ListenAndServe(":8080", nil)

}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Zhiltsov Danila Nikitich\n")
	// make GET request
	response, error := http.Get("https://api.ipify.org/?format=text")
	if error != nil {
		fmt.Println(error)
	}

	// read response body
	body, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	response.Body.Close()

	// print response body СКРЫТые 2 и 3 октета айпишника

	regex := regexp.MustCompile("[.][0-9][0-9][0-9][.][0-9][0-9][0-9][.]")
	out := regex.ReplaceAllString(string(body), ".***.***.")
	fmt.Fprint(w, "IP-address: ", out, "\n")
	now := time.Now()
	fmt.Fprintf(w, now.Format("2006-01-02 3:4:5 pm"))
}

func irlIP(w http.ResponseWriter, r *http.Request) {
	rsp, err := http.Get("https://api.ipify.org/?format=text")
	if err != nil {
		fmt.Println(err)
	}

	// read response body
	qwerty, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// close response body
	rsp.Body.Close()

	// print response body все открытое

	fmt.Fprintf(w, "%s", qwerty)

}
