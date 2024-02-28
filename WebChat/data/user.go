package data

import (
	"log"
	"time"
)

type User struct {
	Uuid       string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}

func (user *User) CreateUser() (error error) {
	statement := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer stmt.Close()

	if err != nil {
		log.Fatalln(err)
		return
	}
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&user.Uuid, &user.Created_at)

	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = $1", email).Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.Created_at)
	return

}
