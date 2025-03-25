package services

import (
	"app/src/config"
	"app/src/types"
	"log"
)

var userServiceLogger = log.New(log.Writer(), "[UserService] ", log.LstdFlags)

func GetUsers() []types.User  {
	db := config.GetDB()
	rows, err := db.Query("select * from users")
	if err != nil {
		println(err)
	}
	defer rows.Close()
	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func CreateUser(user types.User) error {
	db := config.GetDB()
	stmt, err := db.Prepare("insert into users(name, email) values(?, ?)")
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	userServiceLogger.Println("User created successfully!")
	return nil
}

func EditUser(user types.User) error {
	db := config.GetDB()
	stmt, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.ID)
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	userServiceLogger.Println("User updated successfully!")
	return nil
}

func DeleteUser(id string) error {
	db := config.GetDB()
	stmt, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		userServiceLogger.Println(err)
		return err
	}
	userServiceLogger.Println("User deleted successfully!")
	return nil
}