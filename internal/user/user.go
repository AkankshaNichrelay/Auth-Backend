package user

import "time"

// User ...
type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	isActive   bool
	lastActive time.Time
}
