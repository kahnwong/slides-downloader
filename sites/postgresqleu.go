package sites

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func PostgresqlEu(page string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("section.page div.container ul li a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://www.postgresql.eu/events/%s/sessions/%s", page, e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to visit talk url: %s", talkUrl)
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("ul.session-links li a.slides", func(e *colly.HTMLElement) {
		slug := e.Attr("href")
		spider.AppendLineToFile(page, fmt.Sprintf("https://www.postgresql.eu%s", slug))
		fmt.Print(".") // for progress bar
	})

	//// start spider
	url := fmt.Sprintf("https://www.postgresql.eu/events/%s/sessions/", page)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
