package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetHeaders(r *http.Request) {
	defaultHeaders := map[string]string{

		"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:69.0) Gecko/20100101 Firefox/69.0",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"Upgrade-Insecure-Requests": "1",
	}

	for key, value := range defaultHeaders {
		r.Header.Set(key, value)
	}

}

func GetDomain(s string) string {

	u, err := url.Parse(s)

	if err != nil {
		log.Println("Failed to parse", s)
		return ""
	}

	return u.Host
}

func PrintResponseBody(res *http.Response) {

	text, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	body := string(text)
	fmt.Println(body)
}
