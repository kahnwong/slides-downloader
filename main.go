package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

func download(event string) {
	// create download dir
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Fatal().Err(err)
	}

	c := colly.NewCollector(
	//colly.AllowedDomains("shed.com")
	)

	// set headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")
	})

	// rate limiting
	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		RandomDelay: 8 * time.Second,
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
		text := e.Attr("href")

		// open file
		f, err := os.OpenFile(fmt.Sprintf("data/%s.txt", event), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// write to file
		if _, err = f.WriteString(fmt.Sprintf("%s\n", text)); err != nil {
			panic(err)
		}

		fmt.Print(".") // for progress bar
	})

	// start spider
	url := fmt.Sprintf("https://%s.sched.com/overview", event)
	fmt.Printf("Start crawling %s\n", url)

	err = c.Visit(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to visit talk overview")
	}

	fmt.Print("\n")
}

func main() {
	events := []string{
		"spiffespiredayna20",
	}
	for _, event := range events {
		download(event)
	}
}
