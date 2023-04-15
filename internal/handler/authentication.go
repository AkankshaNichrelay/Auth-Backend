package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AkankshaNichrelay/Auth-Backend/internal/session"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/user"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

func CleanSessions() {
	// range over sessions
	// check if last active more than session duration
	// if yes then delete the session and expire the cookie

}

func (h *Handler) IsLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session-cookie")
	if err != nil {
		return false
	}
	fmt.Println("Cookie read:", c)
	_, ok := h.auth.DBSessions[c.Value]
	if !ok {
		return false
	}
	return true
}

func (h *Handler) IsRegistered(userEmail string) bool {
	_, ok := h.auth.DBUsers[userEmail]
	if !ok {
		return false
	}
	return true
}

// RegisterUser registers a new user request
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside registerUser")
	// read the user values
	user := user.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("registerUser Err while Decoding. err:", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Internal Server Error")
		return
	}
	fmt.Println(user)

	// check if already registered
	found := h.IsRegistered(user.Email)
	if found {
		fmt.Println("Already registered")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "User already registered. Continue to login")
		return
	}

	// check if user email is unique - unique in D
	// encrypt password
	// store user in DB
	h.auth.DBUsers[user.Email] = user
	render.JSON(w, r, "User registered")
}

// LoginUser verifies if user is registered and returns a session for login
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	// read login email and password
	usr := user.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println("registerUser Err while Decoding. err:", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "Internal Server Error")
		return
	}
	fmt.Println(usr)

	// check if user exists
	found := h.IsRegistered(usr.Email)
	if !found {
		fmt.Println("User not registered")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "User not found. Please Register to continue")
		return
	}

	// check if already logged in
	loggedin := h.IsLoggedIn(r)
	if loggedin {
		// TODO: update the session timestamp
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "You are already logged in")
		return
	}

	// compare password

	// if not logged in then create session
	uuid := uuid.NewV4()
	c := http.Cookie{
		Name:  "session-cookie",
		Value: uuid.String(),
		// Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)

	// user already exists, get user object
	usr = h.auth.DBUsers[usr.Email]

	sess := session.Session{
		SessionId: uuid,
		UserId:    usr.Id,
	}
	h.auth.DBSessions[c.Value] = sess
	render.JSON(w, r, "User logged in successfully.")

}

// LogoutUser logs out a currently logged in user
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {

	// check if is logged in
	loggedin := h.IsLoggedIn(r)
	if !loggedin {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "You are already logged out")
		return
	}

	// set new cookie to delete
	c, err := r.Cookie("session-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "session-cookie",
			Value: "new-test",
			// Secure:   true,
			HttpOnly: true,
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println("Cookie read:", c)

	// delete session entry from db
	delete(h.auth.DBSessions, c.Value)

	// delete cookie -> maxage = -1
	c.MaxAge = -1 // this cookie is now expired
	http.SetCookie(w, c)

}
