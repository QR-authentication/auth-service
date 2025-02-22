package model

type User struct {
	UUID         string `db:"id"`
	CardNumber   string `db:"card_number"`
	Name         string `db:"name"`
	Surname      string `db:"surname"`
	PasswordHash string `db:"password"`
}
