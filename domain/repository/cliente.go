package repository

import (
	"auth-users-aws/domain/entity"
	database "auth-users-aws/services/database"
	"errors"
	"fmt"
)

type client_db struct {
	id         int
	nm_cliente string
	ds_email   string
	cd_cpf     string
}

func GetClienteByCPF(cpf string) (entity.Cliente, error) {

	rows, err := database.Select("cliente", "cd_cpf", cpf)

	var clienteList []entity.Cliente

	for rows.Next() {
		var id int
		var nm_cliente string
		var ds_email string
		var cd_cpf string

		rows.Scan(&id, &nm_cliente, &ds_email, &cd_cpf)

		var cliente entity.Cliente
		cliente.Id = id
		cliente.CPF = cd_cpf
		cliente.Email = ds_email
		cliente.Name = nm_cliente

		fmt.Println("Cliente: ", cliente)

		clienteList = append(clienteList, cliente)
	}

	fmt.Println("ClientList: ", clienteList)

	fmt.Println("ClientLen: ", len(clienteList))

	if len(clienteList) > 0 {
		return clienteList[0], err
	} else {
		return entity.Cliente{}, errors.New("Cliente Not Found")
	}
}

func GetAllClientes() ([]entity.Cliente, error) {

	fmt.Println("ClientList 00000")

	rows, err := database.Select("cliente", "", "")

	var clienteList []entity.Cliente

	fmt.Println("ClientList")
	for rows.Next() {
		var cl client_db
		var cliente entity.Cliente

		rows.Scan(&cl.id, &cl.nm_cliente, &cl.ds_email, &cl.cd_cpf)

		cliente.Id = cl.id
		cliente.CPF = cl.cd_cpf
		cliente.Email = cl.ds_email
		cliente.Name = cl.nm_cliente

		clienteList = append(clienteList, cliente)
	}

	fmt.Println("ClientList", clienteList)

	return clienteList, err
}
