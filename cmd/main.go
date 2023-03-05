package main

import (
	"fmt"
	"os"

	"github.com/brodiep21/4090/search"
	// "github.com/gocolly/colly"
	// "net/smtp"
)

func main() {

	pass := os.Getenv("PSWRD")

	err := search.SearchNewEgg(pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
