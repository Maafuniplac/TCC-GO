package smtp

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail() {
	// Sender data.
	from := "marco.antonio.alvesf@gmail.com"
	password := "maafdrums1"

	// Receiver email address.
	to := []string{
		"marco@uniplaclages.edu.br",
		"tt.sartor@gmail.com",
	}

	log.Println(from, password)
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Teste email enviado via aplicação GOLANG.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Enviado Corretamente!")
}
