package search

import (
	"fmt"

	"github.com/brodiep21/4090/internal/vcard"
	"github.com/gocolly/colly"
)

func SearchNewEgg() {
	Vcards := []*vcard.Vcard{}

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
	//link associated with the card
	c.OnHTML(".item-container", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		v := vcard.New(0, link, true)

		Vcards = append(Vcards, v)
	})
	//price scrape
	count := 0
	c.OnHTML(".price-current", func(h *colly.HTMLElement) {

		span := h.DOM
		price := span.Find("strong").Text()

		cost, err := PriceConversion(price)
		if err != nil {
			return
		}
		Vcards[count].Price = cost
		count++
	})

	// err := internal.Mailinfo()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(Vcards[1:])
	counter := 0
	c.OnHTML(".item-promo", func(r *colly.HTMLElement) {
		if r.Text == "OUT OF STOCK" {
			counter++
		}
		var stock = func() {
			fmt.Println(count)
			for i := counter; i >= 0; i-- {
				Vcards[count-(i+1)].Stock = false
			}
		}
		defer stock()
	})

	c.Visit("https://www.newegg.com/p/pl?N=100007709%20601408874")

}
