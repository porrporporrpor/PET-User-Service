package model

import (
	"gitlab.com/pplayground/pet_tracking/user-service/interfaces"
	"time"
)

type User struct {
	Id			int		`json:"id,omitempty"`
	Username	string		`json:"username"`
	Password	string		`json:"password,omitempty"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Role		string		`json:"role"`
	CreatedAt	string	`json:"CreatedAt,omitempty"`
	UpdatedAt	string	`json"UpdatedAt,omitempty"`
	groupId		int		`json:"groupId"`

}

func NewUser(user User) *User {
	return &User{
		Id: user.Id,
		Username: user.Username,
		Password: interfaces.GeneratePwd(user.Password),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		groupId: user.groupId,
	}
}

func EditUser(user User) *User {
	return &User{
		Id: user.Id,
		Username: user.Username,
		Password: interfaces.GeneratePwd(user.Password),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Role: user.Role,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		groupId: user.groupId,
	}
}
