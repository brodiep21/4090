package main

import (
	// "fmt"

	"github.com/brodiep21/4090/internal"
	"github.com/brodiep21/4090/search"
	// "github.com/gocolly/colly"
	// "net/smtp"
)

type Vcard struct {
	Name  string
	Price string
	Link  string
}

func main() {
	search.SearchNewEgg()
	err := internal.Mailinfo()
	if err != nil {
		panic(err)
	}
}
