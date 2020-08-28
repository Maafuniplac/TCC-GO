package models

import (
	"log"

	"github.com/db"
)

//Aqui cria a estrutura, conjunto de variaveis
type Carro struct {
	Id        int
	Modelo    string
	Placa     string
	Renavam   int
	IdCliente int
}

//Função para buscar todos os carros e mostra na tela
func ExibeTodosOsCarros(idCliente string) []Carro {
	db := db.ConectaComDB()

	SelectDeTodosOsCarros, err := db.Query("SELECT * FROM VEICULOS WHERE ID_CLIENTE=$1 ORDER BY ID ASC", idCliente)
	if err != nil {
		panic(err.Error())
	}

	d := Carro{}
	carros := []Carro{}

	for SelectDeTodosOsCarros.Next() {
		var id, renavam, id_cliente int
		var modelo, placa string

		err = SelectDeTodosOsCarros.Scan(&id, &modelo, &placa, &renavam, &id_cliente)
		if err != nil {
			panic(err.Error())
		}

		d.Id = id
		d.Modelo = modelo
		d.Placa = placa
		d.Renavam = renavam
		d.IdCliente = id_cliente

		carros = append(carros, d)
		log.Println(carros)
	}
	defer db.Close()

	return carros
}

//Função para criar novo cliente
func CriaNovoCarro(modelo, placa string, renavam int, idcliente int) {
	db := db.ConectaComDB()

	insereDadosDb, err := db.Prepare("INSERT INTO VEICULOS(MODELO, PLACA, RENAVAM, ID_CLIENTE) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosDb.Exec(modelo, placa, renavam, idcliente)
	defer db.Close()
}

//Função para deletar o cliente
func ExcluirVeiculo(id string) {
	db := db.ConectaComDB()

	deletarOCarro, err := db.Prepare("DELETE FROM VEICULOS WHERE ID=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOCarro.Exec(id)

	defer db.Close()
}

//Função para editar o cliente
func EditaCarro(id string) Carro {
	db := db.ConectaComDB()

	carroDoBanco, err := db.Query("SELECT * FROM VEICULOS WHERE ID=$1", id)
	if err != nil {
		panic(err.Error())
	}

	carroParaAtualizar := Carro{}

	for carroDoBanco.Next() {
		var id, renavam, idcliente int
		var placa, modelo string

		err = carroDoBanco.Scan(&id, &modelo, &placa, &renavam, &idcliente)
		if err != nil {
			panic(err.Error())
		}

		carroParaAtualizar.Id = id
		carroParaAtualizar.Modelo = modelo
		carroParaAtualizar.Placa = placa
		carroParaAtualizar.Renavam = renavam
		carroParaAtualizar.IdCliente = idcliente
	}
	defer db.Close()
	return carroParaAtualizar
}

//Função para atualizar o cliente
func AtualizaVeic(id int, modelo, placa string, renavam int) {
	db := db.ConectaComDB()

	atualizaVeic, err := db.Prepare("UPDATE VEICULOS SET MODELO=$1, PLACA=$2, RENAVAM=$3 WHERE ID=$4")
	if err != nil {
		log.Println(id, modelo, placa, renavam)
		panic(err.Error())
	}
	atualizaVeic.Exec(modelo, placa, renavam, id)
	defer db.Close()
}
