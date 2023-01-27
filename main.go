package main

import (
	"fmt"

	"github.com/gocolly/colly"
	// "net/smtp"
)

type Vcard struct {
	Name  string
	Price string
	Link  string
}

func main() {

	c := colly.NewCollector(colly.AllowedDomains("www.newegg.com"))

	// set the header to be a typical user-agent, in case the site doesn't like them scrapes.
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})
	//checking response just because
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})
	//kill main if problems arise
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})
	//NEWEGG SCRAPE
	c.OnHTML(".price-current", func(h *colly.HTMLElement) {
		span := h.DOM
		price := span.Find("strong").Text()

		fmt.Println(price)
	})

	c.OnHTML(".item-promo", func(r *colly.HTMLElement) {
		stock := r.Text
		fmt.Println(stock)
	})
	c.OnHTML(".item-container", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		fmt.Println(link)
		// links := make([]string, 0)
		// links = append(links, link[1:]...)
		// fmt.Println(links)
	})
	c.Visit("https://www.newegg.com/p/pl?N=100007709%20601408874")

	// c.OnHTML(".section-title", func(h *colly.HTMLElement) {
	// 	words := h.ChildText("h2.section-title-text.font-xxxl.text-brandblue")
	// 	// fmt.Println("blahblah")
	// 	fmt.Println(words)
	// })

}
