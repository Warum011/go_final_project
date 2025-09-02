package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const createTable = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(255) NOT NULL DEFAULT "",
	comment TEXT NOT NULL DEFAULT "",
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);
`

const createIndex = `
CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) {
	_, err := os.Stat(dbFile)
	install := os.IsNotExist(err)

	DB, err = sql.Open("sqlite", dbFile) // имя драйвера — "sqlite"
	if err != nil {
		log.Fatal("database opening error: ", err.Error())
	}

	if install {
		if _, err := DB.Exec(createTable); err != nil {
			log.Fatal("failed to init database (table): ", err.Error())
		}
		if _, err := DB.Exec(createIndex); err != nil {
			log.Fatal("failed to init database (index): ", err.Error())
		}
	}
}
