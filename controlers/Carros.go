package controlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/models"
)

// Variavel temp busca todos os arquivos html dentro da pasta tmpl
var tempp = template.Must(template.ParseGlob("tmpl/*.html"))

// funcao para abrir tela de veiculos
func Veiculos(w http.ResponseWriter, r *http.Request) {
	idcliente := r.FormValue("idcliente")

	todosOsVeiculos := models.ExibeTodosOsCarros(idcliente)
	temp.ExecuteTemplate(w, "Veiculos", todosOsVeiculos)
}

// funcao para abrir tela de veiculos
func N_Car(w http.ResponseWriter, r *http.Request) {
	idcliente := r.FormValue("idcliente")
	temp.ExecuteTemplate(w, "N_Car", idcliente)
}

//Função para inserir novo veiculos
func InserirVeiculo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		modelo := r.FormValue("modelo")
		placa := r.FormValue("placa")
		renavam := r.FormValue("renavam")
		idcliente := r.FormValue("idcliente")

		renavamConvertidoParaInt, err := strconv.Atoi(renavam)
		if err != nil {
			log.Println("Erro na conversão do Renavam;", err)
		}

		idClienteConvertidoParaInt, err := strconv.Atoi(idcliente)
		if err != nil {
			log.Println("Erro na conversão do Id_Cliente;", err)
		}

		models.CriaNovoCarro(modelo, placa, renavamConvertidoParaInt, idClienteConvertidoParaInt)

		http.Redirect(w, r, "/", 301)
	}
}

//Função para deletar veiculos
func Delete_Car(w http.ResponseWriter, r *http.Request) {
	idDoCarro := r.URL.Query().Get("idveiculo")
	models.ExcluirVeiculo(idDoCarro)
	http.Redirect(w, r, "/", 301)
}

//Função para editar veiculos
func Edit_Car(w http.ResponseWriter, r *http.Request) {
	idDoCarro := r.URL.Query().Get("idveiculo")
	carro := models.EditaCarro(idDoCarro)
	temp.ExecuteTemplate(w, "E_Car", carro)

	http.Redirect(w, r, "/veic", 301)
}

//Função para update do veiculos
func Update_Car(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("idveiculo")
		modelo := r.FormValue("modelo")
		placa := r.FormValue("placa")
		renavam := r.FormValue("renavam")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para Int:", err)
		}

		renavamConvertidoParaInt, err := strconv.Atoi(renavam)
		if err != nil {
			log.Println("Erro na conversão do Renavam;", err)
		}

		models.AtualizaVeic(idConvertidaParaInt, modelo, placa, renavamConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
