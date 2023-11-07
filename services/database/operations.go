package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Select(tablename string, column string, value string) (*sql.Rows, error) {

	db := Connection()

	defer db.Close() // Certifique-se de fechar a conexão quando não for mais necessária

	// Crie a consulta SQL
	query := "SELECT * FROM cliente"

	var rows *sql.Rows
	var err error

	if value != "" {
		query = fmt.Sprintf("%s WHERE %s = $1", query, column)
		rows, err = db.Query(query, value)

	} else {
		rows, err = db.Query(query)

	}

	if err != nil {
		fmt.Println("Erro ao executar a consulta:", err)
	}
	//defer rows.Close() // Certifique-se de fechar as linhas (rows) quando não forem mais necessárias

	// Leia os resultados da consulta
	// for rows.Next() {
	// 	var id int
	// 	var nm_cliente string
	// 	var ds_email string
	// 	var cd_cpf string
	// 	if err := rows.Scan(&id, &nm_cliente, &ds_email, &cd_cpf); err != nil {
	// 		fmt.Println("Erro ao ler os resultados:", err)
	// 	}
	// 	fmt.Println("result: ", id, nm_cliente, ds_email, cd_cpf)
	// }

	// if err := rows.Err(); err != nil {
	// 	fmt.Println("Erro nos resultados:", err)
	// }

	// conn.Close()

	return rows, err
}

func ValidTableCliente() {
	conn := Connection()
	tableName := "cliente"

	query := fmt.Sprintf("SELECT to_regclass('%s')", tableName)
	var result string
	conn.QueryRow(query).Scan(&result)

	if result == "" {
		createTableQuery := `
		CREATE TABLE table_name(  
			id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			nm_ciente VARCHAR(255)
			ds_email VARCHAR(255)
			cd_cpf VARCHAR(255)
		)
		`
		_, err := conn.Exec(createTableQuery)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Tabela %s criada com sucesso.\n", tableName)
	} else {
		fmt.Printf("A tabela %s já existe.\n", tableName)
	}

	conn.Close()

}
