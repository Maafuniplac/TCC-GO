package models

import (
	"fmt"
	"log"
	"net/smtp"

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
func EnviaEmail(email string, msg []byte) {
	// Sender data.
	from := "marco.antonio.alvesf@gmail.com"
	password := "maafdrums1"

	log.Println("log 1: ", email, msg)
	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)

	log.Println(to, message)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	log.Println("auth ", auth)
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	log.Println(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Enviado Corretamente!")
}
