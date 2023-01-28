package search_test

import (
	"github.com/brodiep21/4090/search"
	"testing"
)

func Test_priceConversion(t *testing.T) {
	t.Run("Returns value w/ basic string", func(t *testing.T) {
		price := "1,800"

		want := 1800
		got, err := search.PriceConversion(price)
		if err != nil {
			t.Error(err)
		}

		if want != got {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Removes multiple symbols for value", func(t *testing.T) {
		price := "$$,1,800..."

		want := 1800
		got, err := search.PriceConversion(price)
		if err != nil {
			t.Error(err)
		}

		if want != got {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("Error for incorrect value", func(t *testing.T) {
		t.Skip()
		price := "badvalue"

		want := 0
		got, err := search.PriceConversion(price)
		if err != nil {
			t.Error(err)
		}

		if want != got {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}
