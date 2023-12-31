package src

import (
	"fmt"

	"github.com/blackflagsoftware/agenda/config"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
)

type (
	Mysql struct{}
)

func (m *Mysql) ConnectDB(c Connection, rootDB bool) (*sqlx.DB, error) {
	var db *sqlx.DB
	dbName := c.DB
	user := c.User
	pwd := c.Pwd
	if rootDB {
		dbName = "mysql"
		if c.AdminUser != "" {
			user = c.AdminUser
		}
		if c.AdminPwd != "" {
			pwd = c.AdminPwd
		}
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, pwd, c.Host, dbName)
	if pwd == "" {
		conn = fmt.Sprintf("%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, c.Host, dbName)
	}
	db, errOpen := sqlx.Open("mysql", conn)
	if errOpen != nil {
		return db, fmt.Errorf("ConnectDB[mysql]: unable to open DB %s****; %s", conn[:6], errOpen)
	}
	return db, nil
}

func (m *Mysql) CheckDB(db *sqlx.DB, dbName string) error {
	checkSql := fmt.Sprintf("SELECT EXISTS(SELECT schema_name FROM information_schema.schemata WHERE schema_name = lower('%s'))", dbName)
	exists := false
	err := db.Get(&exists, checkSql)
	if err != nil {
		return fmt.Errorf("CheckDB[mysql]: unable to check for existing database; %s", err)
	}
	if !exists {
		createSql := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if _, err := db.Exec(createSql); err != nil {
			return fmt.Errorf("CheckDB[mysql]: unable to create database; %s", err)
		}
	}
	return nil
}

func (m *Mysql) CheckTable(db *sqlx.DB) error {
	checkSql := "SELECT EXISTS(SELECT table_name FROM information_schema.tables WHERE table_name = 'migration')"
	exists := false
	err := db.Get(&exists, checkSql)
	if err != nil {
		return fmt.Errorf("CheckTable[mysql]: unable to check for existing table; %s", err)
	}
	if !exists {
		createSql := "CREATE TABLE migration (id serial, file_name varchar(100) NOT NULL)"
		if _, err := db.Exec(createSql); err != nil {
			return fmt.Errorf("CheckTable[mysql]: unable to create table; %s", err)
		}
	}
	return nil
}

func (m *Mysql) LockTable(db *sqlx.DB) bool {
	sqlLock := fmt.Sprintf("SELECT GET_LOCK('%s', -1)", config.AppName)
	success := false
	errLock := db.Get(&success, sqlLock)
	if errLock != nil {
		fmt.Printf("LockTable[mysql]: unable to lock resource; %s", errLock)
		return false
	}
	return true
}

func (m *Mysql) UnlockTable(db *sqlx.DB) error {
	sqlUnlock := fmt.Sprintf("SELECT RELEASE_LOCK('%s')", config.AppName)
	success := null.Bool{}
	errUnlock := db.Get(&success, sqlUnlock)
	if errUnlock != nil {
		return fmt.Errorf("LockTable[mysql]: unable to unlock; %s", errUnlock)
	}
	if !success.Valid {
		// can be set as null
		return nil
	}
	if !success.Bool {
		return fmt.Errorf("LockTable[mysql]: unable to lock with no errors")
	}
	return nil
}
