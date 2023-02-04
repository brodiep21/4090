package vcard_test

import (
	"reflect"
	"testing"

	"github.com/brodiep21/4090/internal/vcard"
)

func TestVcard(t *testing.T) {
	t.Run("add a vcard", func(t *testing.T) {
		got := vcard.New(1500, "some link", false)

		want := &vcard.Vcard{
			Price: 1500,
			Link:  "some link",
			Stock: false,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
	t.Run("add 2 cards", func(t *testing.T) {
		got := vcard.New(1500, "some link", false)
		got2 := vcard.New(1700, "another link", true)

		want := &vcard.Vcard{
			Price: 1500,
			Link:  "some link",
			Stock: false,
		}
		want2 := &vcard.Vcard{
			Price: 1700,
			Link:  "another link",
			Stock: true,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %v, got %v", want, got)
		}

		if !reflect.DeepEqual(want2, got2) {
			t.Errorf("wanted %v, got %v", want2, got2)
		}

	})
}
