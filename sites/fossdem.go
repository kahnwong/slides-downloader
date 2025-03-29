package sites

import (
	"fmt"
	"path"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func FossdemSpider(event string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("table.table.table-striped.table-bordered.table-condensed tbody tr td a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://fosdem.org%s", e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("html body.schedule-event div#main ul.event-attachments.unstyled li a", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		eventParts := strings.Split(event, "/")
		eventSlug := fmt.Sprintf("fossdem-%s", path.Base(eventParts[len(eventParts)-4]))

		spider.AppendLineToFile(eventSlug, fmt.Sprintf("https://fosdem.org%s", text))
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
