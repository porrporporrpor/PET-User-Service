package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/model"
	"log"
)

func GetUsers(db *sql.DB) ([]map[string]interface{}, error) {
	tableName := "users"
	query := fmt.Sprintf(`SELECT id, username, firstName, lastName, email, role FROM %v`, tableName)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Cannot get users information.")
		log.Println("error >>> ", err)
		return nil, errors.New("cannot get users information")
	}
	defer results.Close()

	var user model.User
	var userList []map[string]interface{}
	users := make(map[string]interface{})
	for results.Next() {
		err = results.Scan(
			&user.Id,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Role)
		if err != nil {
			log.Println("This userId is not existed in database.")
			log.Println("error >>> ", err)
			return nil, errors.New("this userId is not existed in database")
		}
		//log.Println(user)
		users["id"] = user.Id
		users["username"] = user.Username
		users["firstName"] = user.FirstName
		users["lastName"] = user.LastName
		users["email"] = user.Email
		users["role"] = user.Role
		userList = append(userList, users)

	}
	log.Println("Query all user information successfully.")
	return userList, nil
}

func GetUser(db *sql.DB, userId int64) (map[string]interface{}, error) {
	results := db.QueryRow(`SELECT id, username, firstName, lastName, email, role FROM users WHERE id=?`, userId)

	var user model.User
	users := make(map[string]interface{})
	err := results.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Role)
	if err != nil {
		log.Println("This userId is not existed in database.")
		log.Println("error >>> ", err)
		return nil, errors.New("this userId is not existed in database")
	}
	users["id"] = user.Id
	users["username"] = user.Username
	users["firstName"] = user.FirstName
	users["lastName"] = user.LastName
	users["email"] = user.Email
	users["role"] = user.Role
	log.Println("Query user information successfully.")
	return users, nil
}

func CreateUser(db *sql.DB, user *model.User) error {
	_, err := db.Exec(`
				INSERT INTO users (username, password, firstName, lastName, email, role, createdAt)
				VALUES (?, ?, ?, ?, ?, "user", ?)`,
				user.Username,
				user.Password,
				user.FirstName,
				user.LastName,
				user.Email,
				user.CreatedAt)
	if err != nil {
		log.Println("Cannot insert new user into database.")
		log.Println("error >>> ", err)
		if mysqlError, _ := err.(*mysql.MySQLError); mysqlError.Number == 1062 {
				return errors.New("duplicate username")
			}
		return errors.New("cannot insert new user into database")
		}
	log.Println("Create user successfully.")
	return nil
}

func UpdateUser(db *sql.DB, user *model.User) error {
	update, err := db.Exec(`
							UPDATE users 
							SET username=?, 
							password=?, 
							firstName=?, 
							lastName=?,
							email=?,
							updatedAt=? 
							WHERE id=?`,
							user.Username,
							user.Password,
							user.FirstName,
							user.LastName,
							user.Email,
							user.UpdatedAt,
							user.Id)
	if rowAffected, _ := update.RowsAffected(); rowAffected == 0 {
		return errors.New("this user is not existed in database")
	}
	if err != nil {
		log.Println("Cannot update user information.")
		log.Println("error >>> ", err)
		return errors.New("cannot update user information")
	}
	log.Println("Update user information successfully.")
	return nil
}

func DeleteUser(db *sql.DB, userId int64) error {
	delete, err := db.Exec(`DELETE FROM users WHERE id=?`, userId)
	if rowAffected, _ := delete.RowsAffected(); rowAffected == 0 {
		return errors.New("this user is not existed in database")
	}
	if err != nil {
		log.Println("Cannot delete user information.")
		log.Println("error >>> ", err)
		return errors.New("cannot delete user information")
	}
	log.Println("Delete user information successfully.")
	return nil
}