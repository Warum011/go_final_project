package main

import (
	"os"

	"github.com/Warum011/go_final_project/pkg/db"
	serv "github.com/Warum011/go_final_project/pkg/server"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	db.Init(dbFile)
	defer db.DB.Close()

	serv.StartServer("web")
}
