package vcard

type Vcard struct {
	Price int
	Link  string
	Stock bool
}

func New(price int, link string, stock bool) *Vcard {
	return &Vcard{
		Price: price,
		Link:  link,
		Stock: stock,
	}
}
