package models

type AddUser struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
