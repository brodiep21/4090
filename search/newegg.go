package search

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Vcard struct {
	Price int
	Link  string
	Stock bool
}

func SearchNewEgg() {
	Vcards := []*Vcard{}

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
		// fmt.Println(link)
		v := &Vcard{
			Price: 0,
			Link:  link,
			Stock: true,
		}
		Vcards = append(Vcards, v)
	})
	count := 0
	//price scrape
	c.OnHTML(".price-current", func(h *colly.HTMLElement) {
		span := h.DOM
		price := span.Find("strong").Text()
		cost, err := PriceConversion(price)
		if err != nil {
			return
		}
		Vcards[count].Price = cost
		count++
		fmt.Println(Vcards[0])
	})
	//checking availability of card
	// c.OnHTML(".item-promo", func(r *colly.HTMLElement) {
	// 	stock := r.Text
	// 	fmt.Println(stock)

	// })

	c.Visit("https://www.newegg.com/p/pl?N=100007709%20601408874")
}
