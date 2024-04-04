package main

import (
	"fmt"

	"github.com/9Neechan/AvitoTech-2024/internal/psql"
)

func main() {
	db, err := psql.Connect()
    if err != nil {
        panic(err)
    }
    defer db.Close()
  
    fmt.Println("Successfully connected to DB!")
}