package sites

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func PgEvents(page string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("div li a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://www.pgevents.ca/events/%s/sessions/%s", page, e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to visit talk url: %s", talkUrl)
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("div ul li a", func(e *colly.HTMLElement) {
		text := e.Attr("href")

		if strings.HasSuffix(text, ".pdf") {
			if !strings.HasPrefix(text, "https://") {
				spider.AppendLineToFile(page, fmt.Sprintf("https://www.pgevents.ca%s", text))
			} else {
				spider.AppendLineToFile(page, text)
			}
		}
		fmt.Print(".") // for progress bar
	})

	//// start spider
	url := fmt.Sprintf("https://www.pgevents.ca/events/%s/sessions/", page)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
