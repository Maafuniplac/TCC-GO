package controlers

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func Email(w http.ResponseWriter, r *http.Request) {
	idDoCliente := r.URL.Query().Get("idcliente")
	temp.ExecuteTemplate(w, "Email", idDoCliente)
	log.Println("idDoCliente: ", idDoCliente)
	http.Redirect(w, r, "/", 301)
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	// Sender data.
	from := "marco.antonio.alvesf@gmail.com"
	password := "maafdrums1"

	log.Println(from, password)
	// Receiver email address.
	to := []string{
		"marco@uniplaclages.edu.br",
		"tt.sartor@gmail.com",
	}

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
	http.Redirect(w, r, "/", 301)
}
