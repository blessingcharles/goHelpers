package crawler

import (
	"Scrapper/utils"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func LinkScrapper(BaseUrl string, depth int, baseDomain string, wg *sync.WaitGroup, urlsChan chan string) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if urlCacheInstance.setVisited(BaseUrl) {
		return
	}

	res, err := utils.DoGetReq(BaseUrl, 10)

	if err != nil {
		//fmt.Println("[-]failed to request ", BaseUrl)
		return

	}

	if res.StatusCode == 200 {

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {

			//fmt.Println("[-]failed to parse", BaseUrl)
			return
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			//finding all href and doing dfs via goroutines

			link, exists := s.Attr("href")

			// \n will cause parser errors so we removed it
			link = strings.Replace(link, "\n", "", -1)

			if !exists {
				return
			}

			NewUrl := ResolveUrl(BaseUrl, link)

			//checking if the url is from the same domain or subdomain
			if strings.Contains(utils.GetDomain(NewUrl), baseDomain) {
				urlsChan <- NewUrl

				wg.Add(1)
				go LinkScrapper(NewUrl, depth-1, baseDomain, wg, urlsChan)
			}

		})
	}

}
