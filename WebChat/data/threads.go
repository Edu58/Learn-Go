package data

import (
	"log"
	"time"
)

type Thread struct {
	Uuid      string
	UserId    string
	Title     string
	CreatedAt time.Time
}

type Threads struct {
	Uuid      string
	User      User
	Title     string
	CreatedAt time.Time
}

func GetThreads() (threads []Threads, err error) {
	rows, err := Db.Query("SELECT t.id, u.name, t.title, t.created_at FROM threads t JOIN users u ON u.id = t.user_id ORDER BY t.created_at DESC")

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		conv := Threads{}
		if err = rows.Scan(&conv.Uuid, &conv.User.Name, &conv.Title, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func (thread *Thread) CreateThread() (err error) {
	statement := "INSERT INTO threads (user_id, title) VALUES ($1, $2) RETURNING id, user_id, title, created_at"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(thread.UserId, thread.Title).Scan(&thread.Uuid, &thread.UserId, &thread.Title, &thread.CreatedAt)

	if err != nil {
		log.Println(err)
		return
	}

	return
}
