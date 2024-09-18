package user

import (
	"errors"
	repository "simpl-commerce/repository/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    string    `json:"userID" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (u *User) CheckLogin(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("password does not match")
	}
	return nil
}

func FromRepoUser(repoUser repository.User) User {
	return User{
		UserID:    repoUser.UserID,
		Phone:     repoUser.Phone,
		Name:      repoUser.Name,
		Password:  repoUser.Password,
		CreatedAt: repoUser.CreatedAt,
		UpdatedAt: repoUser.UpdatedAt.Time,
	}
}

func (u *User) ToProfileResp() GetProfileResp {
	return GetProfileResp{
		Name:  u.Name,
		Phone: u.Phone,
	}
}
