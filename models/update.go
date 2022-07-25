package models

import (
	"API_GO_POSTGRESQL/db"
)

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close() // executa depois que a função Insert terminar fecha a conexão

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`,
		todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	// RowsAffected retorna o número de rows alteradas e um erro
	return res.RowsAffected()

}
