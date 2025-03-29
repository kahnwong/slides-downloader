package sites

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func VirusBulletinSpider(page string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("html.no-js body div.container div.row div.col-md-12 table.table.table-bordered tbody tr td.green-room a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://www.virusbulletin.com%s", e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("html.no-js body div.container.m-top-20 div.row div.col-md-9.col-sm-9.col-lg-9 ul li a", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		spider.AppendLineToFile(page, fmt.Sprintf("https://www.virusbulletin.com%s", text))
		fmt.Print(".") // for progress bar
	})

	//// start spider
	url := fmt.Sprintf("https://www.virusbulletin.com/conference/%s/programme/", page)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
