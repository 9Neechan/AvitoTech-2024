package psql

import (
	"database/sql"
	"os"
	"fmt"
	"gopkg.in/yaml.v3"

	_ "github.com/lib/pq"
)

type PsqlParams struct {
	Host string
    Port int
    User string
    Password string
    DBName string
}

func Connect() (*sql.DB, error) {
	psqlInfo, err := GetPsqlParams()
	if err != nil {
        return nil, err
    }

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
        return nil, err
    }

	err = db.Ping()
    if err != nil {
		return nil, err
    }

	return db, nil
}

func GetPsqlParams() (string, error) {
	var psql_params PsqlParams

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(yamlFile, &psql_params)
	if err != nil {
		return "", err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        psql_params.Host, psql_params.Port, psql_params.User, psql_params.Password, psql_params.DBName)

	return psqlInfo, nil
}