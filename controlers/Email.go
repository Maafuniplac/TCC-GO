package controlers

import (
	"log"
	"net/http"

	"github.com/models"
)

//Aqui ele abre o formulario de email e busca os dados do cliente e veiculo pra exibir na tela
func Email(w http.ResponseWriter, r *http.Request) {
	idcliente := r.URL.Query().Get("idcliente")
	buscaEmail := models.BuscaInfoEmail(idcliente)
	log.Println("idcliente: ", idcliente)
	temp.ExecuteTemplate(w, "Email", buscaEmail)
	http.Redirect(w, r, "/", 301)
}

//Aqui busca as informaçoes do formulario de email para montar o email
func SendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		msg := r.FormValue("msg")

		message := []byte(msg)

		log.Println("SendEmail: ", message)

		models.EnviaEmail(email, message)
		http.Redirect(w, r, "/", 301)
	}

}
