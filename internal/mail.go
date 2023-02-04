package internal

import (
	"os"

	"github.com/go-mail/mail"
)

// []Vcard struct
func Mailinfo() error {
	password := os.Getenv("PSWRD")

	m := mail.NewMessage()

	m.SetHeader("From", "bpeif21@gmail.com")
	m.SetHeader("To", "brodiep21@hotmail.com")
	m.SetHeader("Subject", "Found cards!")

	m.SetBody("text/html", "Insert cards here")

	d := mail.NewDialer("smtp.gmail.com", 587, "bpeif21@gmail.com", password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
