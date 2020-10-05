package controlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/models"
)

//....
var temp = template.Must(template.ParseGlob("tmpl/*.html"))
//....
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsClientes := models.BuscaTodosOsClientes()
	temp.ExecuteTemplate(w, "Index", todosOsClientes)
}

// funcao para abrir tela de clientes
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//Função para inserir novo cliente
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		cpfcnpj := r.FormValue("cpfcnpj")
		telefone := r.FormValue("telefone")
		email := r.FormValue("email")

		cpfcnpjConvertidoParaInt, err := strconv.Atoi(cpfcnpj)
		if err != nil {
			log.Println("Erro na conversão do Cpf/Cnpj;", err)
		}

		telefoneConvertidoParaInt, err := strconv.Atoi(telefone)
		if err != nil {
			log.Println("Erro na conversão do telefone;", err)
		}

		models.CriaNovoCliente(nome, cpfcnpjConvertidoParaInt, telefoneConvertidoParaInt, email)
	}
	http.Redirect(w, r, "/", 301)
}

//Função para deletar cliente
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoCliente := r.URL.Query().Get("id")
	models.DeletaCliente(idDoCliente)
	http.Redirect(w, r, "/", 301)
}

//Função para editar cliente
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoCliente := r.URL.Query().Get("idcliente")
	cliente := models.EditaCliente(idDoCliente)
	temp.ExecuteTemplate(w, "Edit", cliente)

	http.Redirect(w, r, "/", 301)
}

//Função para update do cliente
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		cpfcnpj := r.FormValue("cpfcnpj")
		telefone := r.FormValue("telefone")
		email := r.FormValue("email")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para Int:", err)
		}

		cpfcnpjConvertidoParaInt, err := strconv.Atoi(cpfcnpj)
		if err != nil {
			log.Println("Erro na conversão do Cpf/Cnpj para Int:", err)
		}

		telefoneConvertidoParaInt, err := strconv.Atoi(telefone)
		if err != nil {
			log.Println("Erro na conversão do Telefone para Int:", err)
		}

		models.AtualizaCliente(idConvertidaParaInt, nome, cpfcnpjConvertidoParaInt, telefoneConvertidoParaInt, email)
	}

	http.Redirect(w, r, "/", 301)
}
