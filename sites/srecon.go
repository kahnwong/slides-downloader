package sites

import (
	"fmt"
	"path"
	"strings"

	"github.com/gocolly/colly"
	"github.com/kahnwong/slides-downloader/spider"
	"github.com/rs/zerolog/log"
)

func SreconSpider(event string) {
	c := spider.InitSpider()

	// 1st layer - overview
	c.OnHTML("div.list-simple div.sched-container-inner a", func(e *colly.HTMLElement) {
		firstUrl := fmt.Sprintf("https://%s.sched.com/%s", event, e.Attr("href"))
		err := e.Request.Visit(firstUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit first url")
		}

		fmt.Print(".") // for progress bar
	})

	// 2nd layer - download
	c.OnHTML("article span.usenix-schedule-media.slides a", func(e *colly.HTMLElement) {
		// vars
		eventParts := strings.Split(event, "/")
		eventSlug := path.Base(eventParts[len(eventParts)-2])

		// extract
		text := e.Attr("href") // "/conference/srecon24emea/presentation/curtis"

		textParts := strings.Split(text, "/")
		textProcessed := fmt.Sprintf("https://www.usenix.org/system/files/%s_slides-%s.pdf", eventSlug, path.Base(textParts[len(textParts)-1]))

		// open file
		spider.AppendLineToFile(fmt.Sprintf("%s.txt", eventSlug), textProcessed)

		fmt.Print(".") // for progress bar
	})

	// start spider
	url := event
	fmt.Printf("Start crawling %s\n", url)

	err := c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit conference overview")
	}

	fmt.Print("\n")
}
