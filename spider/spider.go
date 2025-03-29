package spider

import (
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
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		RandomDelay: 8 * time.Second,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed setting rate limiting")
	}

	return c
}
