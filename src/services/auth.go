package services

import (
	"app/src/config"
	"app/src/types"
	"errors"
	"log"
)

var authServiceLogger = log.New(log.Writer(), "[AuthService] ", log.LstdFlags)

func Login(user types.User) (types.User, error) {
	db := config.GetDB()
	rows, err := db.Query("select * from users where name = ? and email = ?", user.Name, user.Email)
	if err != nil {
		authServiceLogger.Println(err)
		return types.User{}, err
	}
	defer rows.Close()

	var loginUser types.User
	if isUser := rows.Next(); !isUser {
		return types.User{}, errors.New("user not found")
	}
	err = rows.Scan(&loginUser.ID, &loginUser.Name, &loginUser.Email)
	if err != nil {
		authServiceLogger.Println(err)
		return types.User{}, err
	}

	authServiceLogger.Println("User logged in successfully!")
	return loginUser, nil
}