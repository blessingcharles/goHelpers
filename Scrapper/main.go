package main

import (
	"Scrapper/crawler"
	"Scrapper/utils"
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func main() {
	var concurrency int
	var payload string
	var myUrl string
	var proxy string
	var depth int

	utils.Banner()

	flag.IntVar(&concurrency, "c", 30, "set concurrency [default 30]")
	flag.StringVar(&payload, "paload", "", "set the payload")
	flag.StringVar(&myUrl, "url", "", "specify the url")
	flag.StringVar(&proxy, "proxy", "", "proxy configuration")
	flag.IntVar(&depth, "depth", 2, "depth to scrap urls")

	flag.Parse()

	if proxy != "" {
		//setting default proxy and ignore ssl secure
		proxyUrl, err := url.Parse(proxy)

		if err != nil {
			panic(err)
		}

		http.DefaultTransport = &http.Transport{

			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxyUrl),
		}

		fmt.Println("[+]proxy --> ", proxy)

	} else {
		http.DefaultTransport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           nil,
		}
	}

	os.Mkdir("Output", 0755)
	s := bufio.NewScanner(os.Stdin)
	wg := &sync.WaitGroup{}

	for s.Scan() {

		wg.Add(1)
		go crawler.Crawler(s.Text(), depth, wg)

	}

	wg.Wait()
}
