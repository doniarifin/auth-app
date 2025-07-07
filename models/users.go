package models

import (
	"database/sql"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func FindUserByEmail(db *sql.DB, email string) (*User, error) {
	query := "SELECT id, email, password FROM users WHERE email = @p1"
	user := &User{}

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// not found
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}


func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (email, password) VALUES (@p1, @p2)`
	_, err := db.Exec(query, user.Email, user.Password)
	return err
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
