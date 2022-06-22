package mail

import (
	"gopkg.in/gomail.v2"
	"os"
)

var list []struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}

func Send() error {

	emailSender := os.Getenv("EMAIL_SENDER")
	passSender := os.Getenv("PASS_SENDER")
	host := os.Getenv("STMP_HOST")

	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", "as.lucasalves@gmail.com")

	m.SetHeader("Subject", "Hello GO!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(host, 587, emailSender, passSender)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		println(err.Error())
		return err
	}

	return nil
}
