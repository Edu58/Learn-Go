package data

import (
	"log"
	"time"
)

type Session struct {
	Uuid       string
	Email      string
	UserId     string
	Created_at time.Time
}

// CreateSession creates a new session for the user.
//
// It takes no parameters and returns a Session and an error.
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (email, user_id, created_at) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET email=excluded.email RETURNING id, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Email, user.Uuid, time.Now()).Scan(&session.Uuid, &session.Email, &session.UserId, &session.Created_at)
	return
}

// Session returns the session for the user.
//
// It takes no parameters and returns a Session and an error.
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Uuid).
		Scan(&session.Uuid, &session.Email, &session.UserId, &session.Created_at)
	return
}

// Check checks the session validity.
//
// No parameters.
// Returns a boolean and an error.
func (session *Session) Check() (valid bool, err error) {
	statement := "SELECT id, email, user_id, created_at FROM sessions WHERE id = $1"
	err = Db.QueryRow(statement, session.Uuid).Scan(&session.Uuid, &session.Email, &session.UserId, &session.Created_at)

	if err != nil {
		valid = false
		return
	}

	if session.Uuid != "" {
		valid = true
	}

	return
}

func (session *Session) DeleteSession() {
	statement := "DELETE FROM sessions WHERE id = $1"
	_, err := Db.Exec(statement, session.Uuid)
	if err != nil {
		log.Println(err)
		return
	}
}
