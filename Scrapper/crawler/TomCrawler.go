package crawler

import (
	"Scrapper/utils"
	"fmt"
	"os"
	"sync"
)

type urlCache struct {
	key   sync.Mutex
	store map[string]bool
}

var urlCacheInstance = urlCache{store: make(map[string]bool)}

func (m *urlCache) setVisited(url string) bool {

	// checking if the url already visited by setting mutex locks
	m.key.Lock()
	defer m.key.Unlock()

	if m.store[url] {
		return true
	}
	m.store[url] = true

	return false
}

func Crawler(baseUrl string, depth int, outerWg *sync.WaitGroup) {

	// fmt.Println("[+] Crawler Started on ", baseUrl)

	urlsChan := make(chan string, 100)
	domain := utils.GetDomain(baseUrl)
	wg := &sync.WaitGroup{}

	wg.Add(1)

	output_file := "Output/" + domain + ".txt"
	f, err := os.Create(output_file)
	// fmt.Println("[+]Output --> ", output_file)
	utils.CheckErr(err)

	go func() {
		for url := range urlsChan {
			fmt.Fprintln(f, url)
			fmt.Println(url)
		}
	}()

	LinkScrapper(baseUrl, depth, domain, wg, urlsChan)

	wg.Wait()
	close(urlsChan)
	//crawler routine is finished
	outerWg.Done()

}
