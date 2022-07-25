package models

import "API_GO_POSTGRESQL/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close() // executa depois que a função Insert terminar fecha a conexão

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
