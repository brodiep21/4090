package search

import (
	"strconv"
	"strings"
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
