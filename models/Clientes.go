package models

import (
	"github.com/db"
)

//Aqui cria a estrutura, conjunto de variaveis
type Cliente struct {
	Id       int
	Nome     string
	CpfCnpj  int
	Telefone int
	Email    string
}

//Função para buscar todos os clientes e mostrar na tela
func BuscaTodosOsClientes() []Cliente {
	db := db.ConectaComDB()

	SelectDeTodosOsClientes, err := db.Query("SELECT * FROM CLIENTES ORDER BY ID ASC")
	if err != nil {
		panic(err.Error())
	}

	c := Cliente{}
	clientes := []Cliente{}

	for SelectDeTodosOsClientes.Next() {
		var id, cpfcnpj, telefone int
		var nome, email string

		err = SelectDeTodosOsClientes.Scan(&id, &nome, &cpfcnpj, &telefone, &email)
		if err != nil {
			panic(err.Error())
		}

		c.Id = id
		c.Nome = nome
		c.CpfCnpj = cpfcnpj
		c.Telefone = telefone
		c.Email = email

		clientes = append(clientes, c)
	}
	defer db.Close()
	return clientes
}

//Função para criar novo cliente
func CriaNovoCliente(nome string, cpfcnpj, telefone int, email string) {
	db := db.ConectaComDB()

	insereDadosDb, err := db.Prepare("INSERT INTO CLIENTES(NOME, CPFCNPJ, TELEFONE, EMAIL) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosDb.Exec(nome, cpfcnpj, telefone, email)
	defer db.Close()
}

//Função para deletar o cliente
func DeletaCliente(id string) {
	db := db.ConectaComDB()

	deletarOCliente, err := db.Prepare("DELETE FROM CLIENTES WHERE ID=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOCliente.Exec(id)

	defer db.Close()
}

//Função para editar o cliente
func EditaCliente(id string) Cliente {
	db := db.ConectaComDB()

	clienteDoBanco, err := db.Query("SELECT * FROM CLIENTES WHERE ID=$1", id)
	if err != nil {
		panic(err.Error())
	}

	clienteParaAtualizar := Cliente{}

	for clienteDoBanco.Next() {
		var id, cpfcnpj, telefone int
		var nome, email string

		err = clienteDoBanco.Scan(&id, &nome, &cpfcnpj, &telefone, &email)
		if err != nil {
			panic(err.Error())
		}

		clienteParaAtualizar.Id = id
		clienteParaAtualizar.Nome = nome
		clienteParaAtualizar.CpfCnpj = cpfcnpj
		clienteParaAtualizar.Telefone = telefone
		clienteParaAtualizar.Email = email
	}
	defer db.Close()
	return clienteParaAtualizar
}

//Função para atualizar o cliente
func AtualizaCliente(id int, nome string, cpfcnpj, telefone int, email string) {
	db := db.ConectaComDB()

	AtualizaCliente, err := db.Prepare("UPDATE CLIENTES SET NOME=$1, CPFCNPJ=$2, TELEFONE=$3, EMAIL=$4 WHERE ID=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaCliente.Exec(nome, cpfcnpj, telefone, email, id)
	defer db.Close()
}
