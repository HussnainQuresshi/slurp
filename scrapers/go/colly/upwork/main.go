// Upwork rejects as bot program.

package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly/v2"
)

func main() {
    allowed := "www.upwork.com"
    url := "https://%s/fl/%s"

    // Instantiate default collector
    c := colly.NewCollector(
        // Visit only domains: https://www.upwork.com
        colly.AllowedDomains(allowed),
    )

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Something went wrong:", err)
    })

    c.OnXML("//script", func(e *colly.XMLElement) {

    })
    // Start scraping on https://www.upwork.com
    err := c.Visit(fmt.Sprintf(url, allowed, "johnbampton"))

    if err != nil {
        log.Fatal(err)
    }
}
