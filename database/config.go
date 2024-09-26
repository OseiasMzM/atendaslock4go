package database

import (
	"fmt"
)

func Config() {
	dbConfig := NewDatabaseConfig()
	db, err := ConnectDB(dbConfig)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	// Configurar o timezone para a sessão
	_, err = db.Exec("SET TIMEZONE = 'America/Sao_Paulo'")
	if err != nil {
		panic(err)
	}
	// Testar a configuração
	var now string
	row := db.QueryRow("SELECT NOW()")
	err = row.Scan(&now)
	if err != nil {
		panic(err)
	}

	fmt.Println("Current time:", now)

}
