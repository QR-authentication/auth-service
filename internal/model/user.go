package model

import "time"

type User struct {
	UUID         string    `db:"id"`
	Name         string    `db:"name"`
	Surname      string    `db:"surname"`
	Patronymic   string    `db:"patronymic"`
	Department   string    `db:"department"`
	GroupNumber  string    `db:"group_number"`
	CardNumber   string    `db:"card_number"`
	PasswordHash string    `db:"password"`
	createdAt    time.Time `db:"created_at"`
}
