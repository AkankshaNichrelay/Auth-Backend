package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/AkankshaNichrelay/Auth-Backend/internal/session"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/user"
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

// Close the db connection
func (mysql *Mysql) Close() {
	if mysql.db == nil {
		return
	}
	err := mysql.db.Close()
	if err != nil {
		mysql.log.Println("MySql Close failed, err:", err.Error())
	}
}

// FetchRow
func (mysql *Mysql) FetchRow(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error) {
	row := mysql.db.QueryRowContext(ctx, query, args...)

	// This should ideally be handled by creating a generic struct mapper using reflection
	switch result.(type) {
	case user.User:
		if err := row.Scan(result.(user.User).Id, result.(user.User).Email, result.(user.User).Password); err != nil {
			if err == sql.ErrNoRows {
				return 0, nil
			}
			return 0, fmt.Errorf("FetchRow failed err: %v", err)
		}
	case session.Session:
		if err := row.Scan(result.(session.Session).SessionId, result.(session.Session).UserId); err != nil {
			if err == sql.ErrNoRows {
				return 0, nil
			}
			return 0, fmt.Errorf("FetchRow failed err: %v", err)
		}
	default:
		mysql.log.Println("Type not handled")
	}

	return 0, nil
}

// Exec
func (mysql *Mysql) Exec(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}
