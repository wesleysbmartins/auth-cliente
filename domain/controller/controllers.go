package controllers

import (
	"auth-users-aws/domain/controller/dto"
	"auth-users-aws/domain/usecase"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func GetClientByCpfHandler(w http.ResponseWriter, r *http.Request) {
	var cliente dto.ClienteDTO

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &cliente); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	value, err := usecase.GetClienteByCPFUseCase(cliente.Cpf)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	responseJSON, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Erro ao serializar o cliente", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetAllClientesHandler(w http.ResponseWriter, r *http.Request) {
	value := usecase.GetAllClientesFromCache()

	responseJSON, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Erro ao serializar o cliente", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
