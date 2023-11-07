package dto

type ClienteDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}
