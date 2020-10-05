package models

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/db"
)

//Função para buscar todos os clientes e mostrar na tela
func BuscaInfoEmail(idCliente string) Cliente {
	db := db.ConectaComDB()

	SelectCliente, err := db.Query("SELECT NOME, EMAIL FROM CLIENTES WHERE ID =$1", idCliente)
	if err != nil {
		panic(err.Error())
	}

	clienteParaEmail := Cliente{}

	for SelectCliente.Next() {
		var nome, email string

		err = SelectCliente.Scan(&nome, &email)
		if err != nil {
			panic(err.Error())
		}

		clienteParaEmail.Nome = nome
		clienteParaEmail.Email = email
		log.Println("BuscaInfoEmail Select: ", nome, email)
	}
	defer db.Close()
	log.Println("return clientes: ", clienteParaEmail)
	return clienteParaEmail
}

//Funcao para enviar o email
func EnviaEmail(email string, msg []byte, nome string) {
	// Sender data.
	from := "marco.antonio.alvesf@gmail.com"
	password := "maafdrums1"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)
	cliente := nome
	//log.Println(to, message)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//log.Println("auth ", auth)
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	//log.Println(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Erro no envio do email: ", err)
		return
	}

	//datetime.Format(("02/Jan/2006 15:04:05"))
	//log.Println("datetime: ", datetime)

	InsereEmail(msg, cliente)

	fmt.Println("Email Enviado Corretamente!")
}

//Inserir no banco
func InsereEmail(mensagem []byte, nome string) {
	db := db.ConectaComDB()

	datetime := time.Now().UTC()

	insereDadosDb, err := db.Prepare("INSERT INTO EMAIL(MENSAGEM, CLIENTE, DATA) VALUES($1, $2, $3)")
	log.Println("insereDadosDb: ", insereDadosDb)
	if err != nil {
		panic(err.Error())
	}

	insereDadosDb.Exec(mensagem, nome, datetime)
	defer db.Close()
}
