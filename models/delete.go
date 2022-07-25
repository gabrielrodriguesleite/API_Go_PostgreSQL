package models

import (
	"API_GO_POSTGRESQL/db"
)

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close() // executa depois que a função Insert terminar fecha a conexão

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	// RowsAffected retorna o número de rows alteradas e um erro
	return res.RowsAffected()

}
