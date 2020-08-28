package main

import (
	"log"      // Pacote Log, para registro de erros em log
	"net/http" // Gerencia URLs e Servidor Web

	"github.com/routes"
)

func main() {
	routes.CarregaRotas()
	// Exibe mensagem que o servidor foi iniciado
	log.Println("Server started on: http://localhost:8000")

	// Inicia o servidor na porta 8000
	http.ListenAndServe(":8000", nil)
}
