package auth

import (
	"github.com/AkankshaNichrelay/Auth-Backend/internal/db"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/session"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/user"
)

type Auth struct {
	db         *db.Mysql
	DBUsers    map[string]user.User       //user email, user
	DBSessions map[string]session.Session // session ID, session
}

// New creates new Authenticator instance
func New(db *db.Mysql) *Auth {
	a := Auth{
		db:         db,
		DBUsers:    map[string]user.User{},
		DBSessions: map[string]session.Session{},
	}
	return &a
}
