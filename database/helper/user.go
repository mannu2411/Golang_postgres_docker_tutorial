package helper

import (
	"database/sql"
	"github.com/tejashwikalptaru/tutorial/database"
	"github.com/tejashwikalptaru/tutorial/models"
)

func CreateUser(name, email string) (string, error) {
	// language=SQL
	SQL := `INSERT INTO users(name, email) VALUES ($1, $2) RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, name, email)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func GetUser(userID string) (*models.User, error) {
	// language=SQL
	SQL := `SELECT id, name, email, created_at, archived_at FROM users WHERE archived_at IS NULL AND id = $1`
	var user models.User
	err := database.Tutorial.Get(&user, SQL, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func UpdateUser(name, email, id string) (string, error) {
	//language=SQL
	SQL := `UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, name, email, id)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func DeleteUser(uid string) sql.Result {
	//language=SQL
	SQL := `DELETE FROM users WHERE id=$1 RETURNING id;`

	err, _ := database.Tutorial.Exec(SQL, uid)

	return err
}
