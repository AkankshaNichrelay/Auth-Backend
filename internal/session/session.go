package session

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Session ...
type Session struct {
	SessionId  uuid.UUID `json:"session_id"`
	UserId     int       `json:"user_id"`
	lastActive time.Time
}
