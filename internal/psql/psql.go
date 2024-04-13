package psql

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"gopkg.in/yaml.v3"
	_ "github.com/lib/pq"

	"github.com/9Neechan/AvitoTech-2024/internal/database"
)

var DB *database.Queries = GetDBQueries()

type PsqlParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// host=localhost port=5432 user=postgres password=123 dbname=banners sslmode=disable

func ConnectDB() *sql.DB {
	psqlInfo, err := GetPsqlParams()
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to DB!")

	return db
}

func GetDBQueries() *database.Queries {
	db_conn := ConnectDB()
	dbQueries := database.New(db_conn)

	return dbQueries
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