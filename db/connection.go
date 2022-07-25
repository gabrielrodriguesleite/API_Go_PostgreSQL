package db

import (
	"database/sql"
	"fmt"
	"log"

	// importação do pacote definido dentro do módulo atual API_GO_POSTGRESQL
	"API_GO_POSTGRESQL/configs"

	// este import é necessário mas será removido pela IDE se não tiver o _
	_ "github.com/lib/pq"
)

// abre a conexão com bd

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// string connection
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Printf("Erro na conexão: host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
		panic(err) // (!ATENÇÃO) não usar em produção, derruba a aplicação
	}

	err = conn.Ping() // mesmo que fazer um SELECT 1 pra testar a conexão

	return conn, err
}
