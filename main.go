package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.newegg.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})
	c.Visit("https://www.newegg.com")

	c.OnHTML(".section-title", func(h *colly.HTMLElement) {
		words := h.ChildText("h2.section-title-text.font-xxxl.text-brandblue")
		// fmt.Println("blahblah")
		fmt.Println(words)
	})
}
