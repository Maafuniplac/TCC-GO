package routes

import (
	"net/http"

	"github.com/controlers"
)

// Gerencia as URLs, para onde deve ir
func CarregaRotas() {
	http.HandleFunc("/", controlers.Index)
	http.HandleFunc("/new", controlers.New)
	http.HandleFunc("/insert", controlers.Insert)
	http.HandleFunc("/delete", controlers.Delete)
	http.HandleFunc("/edit", controlers.Edit)
	http.HandleFunc("/update", controlers.Update)
	http.HandleFunc("/veic", controlers.Veiculos)
	http.HandleFunc("/novoCarro", controlers.N_Car)
	http.HandleFunc("/inserirVeiculo", controlers.InserirVeiculo)
	http.HandleFunc("/deleteVeic", controlers.Delete_Car)
	http.HandleFunc("/editar_vei", controlers.Edit_Car)
	http.HandleFunc("/atualizaVeic", controlers.Update_Car)
	http.HandleFunc("/excluiVeic", controlers.Delete_Car)

	http.HandleFunc("/email", controlers.Email)
	http.HandleFunc("/enviaEmail", controlers.SendEmail)
}
