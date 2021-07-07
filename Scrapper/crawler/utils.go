package crawler

import (
	"fmt"
	"log"
	"net/url"
)

func UrlParse(s string) string {

	u, err := url.Parse(s)

	if err != nil {
		log.Println("Failed to parse", s)
	}

	return u.String()

}

// converting relative url to absolute url
func ResolveUrl(baseUrl string, relative string) string {
	u, err := url.Parse(relative)

	if err != nil {
		fmt.Println("failed to parse", err)
	}

	base, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println("failed to parse", err)
	}

	newUrl := base.ResolveReference(u)

	return newUrl.String()
}
