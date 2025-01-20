package models

import (
	"basicapis/db"
	"basicapis/utils"
	"log"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func GetAllUsers() ([]User, error) {
	getUsersQuery := `SELECT * FROM users`

	rows, err := db.DB.Query(getUsersQuery)
	if err != nil {
		log.Fatalf("error in getting users rows", err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			log.Fatalf("error in scanning users rows", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (u User) SaveUser() error {
	saveUserQuery := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(saveUserQuery)
	if err != nil {
		log.Fatalf("error in saving user", err)
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Fatalf("error in hashing passwrod", err)
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		log.Fatalf("error in executing save user query", err)
	}
	userId, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("error in getting last user id", err)
	}
	u.ID = userId
	return err
}

func (u User) ValidateUser() {
	validateUserQuery := `SELECT * FROM users WHERE email = ?`

	row := db.DB.QueryRow(validateUserQuery, u.Email)

	row.Scan()

}
