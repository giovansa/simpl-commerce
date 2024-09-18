package user

import (
	"database/sql"
	"time"
)

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type RegisterUser struct {
	ID       string `db:"id"`
	Phone    string `db:"phone"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type User struct {
	UserID    string       `json:"userID" db:"id"`
	Phone     string       `json:"phone" db:"phone"`
	Name      string       `json:"name" db:"name"`
	Email     string       `json:"email" db:"email"`
	Password  string       `json:"password" db:"password"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updatedAt" db:"updated_at"`
}
