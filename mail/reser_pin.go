package mail

import (
	gomail "gopkg.in/gomail.v2"
)

func ResetMail(email, note, typer, bundle, amount string) {

	body = ""

	msg := gomail.NewMessage()
	msg.SetHeader("From", "info@payton.jitssolutions.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", note)
	msg.SetBody("text/html", body)
	//msg.Attach("pic.jpg")

	n := gomail.NewDialer("payton.jitssolutions.com", 465, "info@payton.jitssolutions.com", "mylovefordogs1$")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}
