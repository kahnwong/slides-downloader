package spider

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

func InitSpider() *colly.Collector {
	c := colly.NewCollector()

	// set headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")
	})

	// rate limiting
	randomDelaySecond := stringToInt(os.Getenv("RANDOM_DELAY_SECOND"))
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: stringToInt(os.Getenv("PARALLELISM")),
		RandomDelay: time.Duration(randomDelaySecond) * time.Second,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed setting rate limiting")
	}

	return c
}

func stringToInt(s string) int {
	vInt, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return int(vInt)
}
