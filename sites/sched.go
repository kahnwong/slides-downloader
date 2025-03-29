package sites

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kahnwong/sched-downloader/spider"
	"github.com/rs/zerolog/log"
)

func SchedSpider(event string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("div.list-simple div.sched-container-inner a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://%s.sched.com/%s", event, e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("a.file-uploaded", func(e *colly.HTMLElement) {
		text := e.Attr("href")
		spider.AppendLineToFile(event, text)
		fmt.Print(".") // for progress bar
	})

	// start spider
	url := fmt.Sprintf("https://%s.sched.com/overview", event)
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}
