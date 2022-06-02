package models

import (
	"github.com/volatiletech/null"
	"time"
)

type User struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email" json:"email"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	ArchivedAt null.Time `db:"archived_at" json:"archivedAt"`
}

const ActiveUser string = "active_user"
