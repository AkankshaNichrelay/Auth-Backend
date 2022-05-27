package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr     string
	Net      string
	User     string
	Password string
	Database string
}

// Mysql database
type Mysql struct {
	config *Config
	log    *log.Logger
	db     *sql.DB
}

// New creates a new MySQL client instance
func New(log *log.Logger, cfg *Config) (*Mysql, error) {
	// TODO: Put these values into config file
	config := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "auth_backend",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("Ping Err", pingErr)
		return nil, pingErr
	}
	log.Println("MySQL Database Connection established.")

	mysql := Mysql{
		log: log,
		db:  db,
	}

	return &mysql, nil
}

// FetchRows
func FetchRows(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error) {
	// TODO
	return 0, nil
}

// Exec
func Exec(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}
