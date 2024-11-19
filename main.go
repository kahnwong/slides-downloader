package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

var (
	event = "pytorch2024"
	date  = "2024-09-19"
)

func main() {
	c := colly.NewCollector(
	//colly.AllowedDomains("shed.com")
	)

	// set headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")
	})

	// rate limiting
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*sched.com.*",
		Parallelism: 4,
		RandomDelay: 1 * time.Second,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed setting rate limiting")
	}

	// 1st layer - overview
	c.OnHTML("div.list-simple div.sched-container-inner a", func(e *colly.HTMLElement) {
		talkUrl := fmt.Sprintf("https://%s.sched.com/%s", event, e.Attr("href"))
		err := e.Request.Visit(talkUrl)
		if err != nil {
			log.Error().Err(err).Msg("Failed to visit talk url")
		}
	})

	// 2nd layer - download
	c.OnHTML("a.file-uploaded", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	// start spider
	url := fmt.Sprintf("https://%s.sched.com/%s/overview", event, date)
	fmt.Printf("Start crawling %s\n", url)

	err = c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}
}
