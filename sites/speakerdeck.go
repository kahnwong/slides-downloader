package sites

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func SpeakerdeckSpider(page string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("html body.sd-app div.sd-main div.container.py-md-4.py-3 div.row.mt-4.mb-4 div.col-12.col-md-6.col-lg-4.mb-5 div.card.deck-preview a.deck-preview-link", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://speakerdeck.com%s", e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("div.container div.row.align-items-center.justify-content-between div.col-md-auto.col-12.py-md-3.pb-3 div.row.justify-content-between.justify-content-md-start.gap-2.gap-md-0 div.row.col-auto.text-white.font-weight-bold div.col-auto.pe-0.pe-lg-2.align-self-center a.text-white", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		spider.AppendLineToFile(page, text)
		fmt.Print(".") // for progress bar
	})

	//// start spider
	url := fmt.Sprintf("https://speakerdeck.com/%s", page)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
