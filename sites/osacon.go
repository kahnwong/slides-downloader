package sites

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func OsaCon(page string) {
	c := spider.InitSpider()

	c.OnHTML("table td a", func(e *colly.HTMLElement) {
		slug := e.Attr("href")

		if strings.HasSuffix(slug, ".pdf") {
			spider.AppendLineToFile(fmt.Sprintf("osacon-%s", page), fmt.Sprintf("https://osacon.io%s", slug))
		}
		fmt.Print(".") // for progress bar
	})

	//// start spider
	url := fmt.Sprintf("https://osacon.io/sessions/%s/", page)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
