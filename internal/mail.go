package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brodiep21/4090/internal/vcard"
	"gopkg.in/mail.v2"
)

// mails over information
func Mailinfo(information []*vcard.Vcard) error {
	s := make([]string, 0)
	password := os.Getenv("PSWRD")

	for i := 0; i < len(information); i++ {
		s = append(s, fmt.Sprintf("%+v \n %+v \n", strconv.Itoa(information[i].Price), information[i].Link))
	}
	m := mail.NewMessage()

	m.SetHeader("From", "bpeif21@gmail.com")
	m.SetHeader("To", "brodiep21@hotmail.com")
	m.SetHeader("Subject", "Found cards!")

	m.SetBody("text/html", strings.Join(s[:], "\n"))
	//verify stmp env var via gpass
	d := mail.NewDialer("smtp.gmail.com", 587, "bpeif21@gmail.com", password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
