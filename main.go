/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"logaggregator/cmd"
	"logaggregator/storage/sqlite"
)

func main() {
	db, err := sqlite.ConnectToSQLite()

	if err != nil {
		log.Fatal("Xatolik sqlite malumotlar bazasiga ulanishda: ", err.Error())
	}
	defer db.Close()

	cmd.SetDBConnection(db)
	cmd.Execute()
}
