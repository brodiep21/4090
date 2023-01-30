package internal

import (
	"fmt"
	"os"

	"github.com/go-mail/mail"
)

func Mailinfo([]Vcard struct) {
	password := os.Getenv("PSWRD")
	fmt.Println(password)
	m := mail.NewMessage()

	m.SetHeader("From", "bpeif21@gmail.com")
	m.SetHeader("To", "brodiep21@hotmail.com")
	m.SetHeader("Subject", "Found cards!")

	m.SetBody("")

	mail.NewDialer("smtp.gmail.com", 587, "bpeif21@gmail.com", password)
}
