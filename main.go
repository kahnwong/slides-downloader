package main

import (
	"os"

	"github.com/kahnwong/sched-downloader/sites"

	"github.com/rs/zerolog/log"
)

func main() {
	var site string
	if len(os.Args) > 1 {
		site = os.Args[1]
	} else {
		log.Fatal().Msg("No site specified")
	}

	if site == "sched" {
		events := []string{
			"spiffespiredayna20",
		}
		for _, event := range events {
			sites.SchedSpider(event)
		}
	}
}

func init() {
	// create download dir
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Fatal().Err(err)
	}
}
