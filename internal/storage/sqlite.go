package storage

import (
	"fmt"
	"os"

	"github.com/blackflagsoftware/agenda/config"
	m "github.com/blackflagsoftware/agenda/internal/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var SqliteDB *sqlx.DB
var sqliteConnection string

func SqliteInit() *sqlx.DB {
	if SqliteDB == nil {
		var err error
		sqliteConnection = GetSqliteConnection()
		SqliteDB, err = sqlx.Open("sqlite3", sqliteConnection)
		if err != nil {
			m.Default.Panicf("Could not connect to the DB host: %s*****; %s", sqliteConnection[6:], err)
		}
		SqliteDB.SetMaxOpenConns(1)
	}
	return SqliteDB
}

func GetSqliteConnection() string {
	if _, err := os.Stat(config.SqlitePath); os.IsNotExist(err) {
		file, err := os.OpenFile(config.SqlitePath, os.O_CREATE, 0644)
		if err != nil {
			m.Default.Panicf("Could not create new .sql file at: %s", config.SqlitePath)
		}
		file.Close()
	}
	return fmt.Sprintf("%s?cache=shared&mode=wrc", config.SqlitePath)
}
