package search

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brodiep21/4090/internal/vcard"
)

func PriceConversion(price string) (int, error) {
	//strip any contexts of outlying values in the string that are not needed for values sake.
	price = strings.Replace(price, ",", "", -1)
	price = strings.Replace(price, "$", "", -1)
	price = strings.Replace(price, ".", "", -1)
	//convert string to int
	cost, err := strconv.Atoi(price)
	if err != nil {
		return 0, err
	}
	return cost, nil
}

func FindValueCards(Vcards []*vcard.Vcard) {
	send := []*vcard.Vcard{}
	for i := 0; i < len(Vcards); i++ {
		if Vcards[i].Price <= 1850 && Vcards[i].Stock {
			send = append(send, Vcards[i])
		}

	}
	fmt.Println(send)
	// internal.Mailinfo(send)
}
