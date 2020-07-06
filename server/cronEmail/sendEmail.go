package cronEmail

import (
	"fmt"
	"net/smtp"
	"os"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// SendEmail : Notify to the user that the price has change
func SendEmail(email, product string) {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("PSW")
	// Receiver email address.
	to := []string{
		email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := fmt.Sprintf("Your product %s has changed", product)

	// Message.

	messageByte := []byte(message)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, messageByte)
	if err != nil {
		fmt.Println(err)
		return
	}
}
