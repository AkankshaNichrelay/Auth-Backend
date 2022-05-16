package auth

import "github.com/AkankshaNichrelay/Auth-Backend/internal/db"

type Auth struct {
	db *db.Mysql
}

// New creates new Authenticator instance
func New(db *db.Mysql) *Auth {
	a := Auth{db: db}
	return &a
}
