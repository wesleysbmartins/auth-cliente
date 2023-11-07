package usecase

import (
	"auth-users-aws/domain/entity"
	"auth-users-aws/domain/repository"
	"auth-users-aws/services/cache"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func GetClienteByCPFUseCase(cpf string) ([]byte, error) {

	for _, value := range cpf {
		tamanho := utf8.RuneLen(value)
		buffer := make([]byte, tamanho)
		utf8.EncodeRune(buffer, value)
		caractere := string(buffer)

		_, err := strconv.Atoi(caractere)
		if err != nil {
			return []byte{}, errors.New("CPF Inválido")
		}
	}

	if len(cpf) != 11 {
		return []byte{}, errors.New("CPF Inválido")
	}

	clienteByte, err := GetClienteByCPFInCache(cpf)

	if len(clienteByte) > 0 && err == nil {
		fmt.Println("Get in CACHE")
		return clienteByte, err
	}

	client, err := repository.GetClienteByCPF(cpf)

	if client.Name != "" && err == nil {
		clientjson, _ := json.Marshal(client)
		cache.Set(client.CPF, clientjson)
		fmt.Println("Get in Disk")
		return clientjson, err
	} else {
		return clienteByte, errors.New("Cliente Não Existe")
	}
}

func GetClienteByCPFInCache(cpf string) ([]byte, error) {

	value, err := cache.Get(cpf)

	if value != "" && err == nil {
		return json.Marshal(value)
	} else {
		return []byte{}, errors.New("Not Found In Redis")
	}
}

func GetClienteByCPFInDisk(cpf string) (entity.Cliente, error) {

	return repository.GetClienteByCPF(cpf)
}

func UpdateCache() {
	cache.Flush()

	clientes, err := repository.GetAllClientes()
	addictions := 0

	if err == nil && len(clientes) > 0 {
		for _, cliente := range clientes {
			clienteJson, _ := json.Marshal(cliente)
			err := cache.Set(cliente.CPF, clienteJson)
			if err != nil {
				panic(err)
			} else {
				addictions += 1
			}
		}

		fmt.Println("Cache Updated")
	}
}

func GetAllClientesFromCache() []string {
	clientes, _ := cache.GetAll()
	return clientes
}
