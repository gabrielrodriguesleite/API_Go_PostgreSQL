package models

import (
	"API_GO_POSTGRESQL/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close() // executa depois que a função Insert terminar fecha a conexão

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// o Scan passa o endereço da variável de retorno criada na definição da função
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return // os valores do retorno já está implícito na definição da função
}
