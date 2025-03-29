package sites

import (
	"fmt"
	"path"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func ScaleSpider(event string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("div.view-id-presentation_page div.view-content span.field-content a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://www.socallinuxexpo.org%s", e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("div.views-field-field-presentation-to-upload div.field-content span.file a", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		eventParts := strings.Split(event, "/")
		eventSlug := path.Base(eventParts[len(eventParts)-2])

		spider.AppendLineToFile(eventSlug, text)
		fmt.Print(".") // for progress bar
	})

	// start spider
	url := event
	fmt.Printf("Start crawling %s\n", event)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
