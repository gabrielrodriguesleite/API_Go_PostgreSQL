package models

import "API_GO_POSTGRESQL/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close() // executa depois que a função Insert terminar fecha a conexão

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	// Next intera os itens lidos do db
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
