package user

import (
	"context"
	"database/sql"
	"errors"
)

func (r *Repository) RegisterUser(ctx context.Context, input RegisterUser) (string, error) {
	res, err := r.Db.ExecContext(ctx, qInsertUser, input.ID, input.Phone, input.Name, input.Password)
	if err != nil {
		return "", err
	}

	if count, _ := res.RowsAffected(); count < 1 {
		return "", errors.New("fail register")
	}

	return input.ID, nil
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	stmt, err := r.Db.PrepareContext(ctx, qGetUserByPhone)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(phone).
		Scan(&user.UserID, &user.Phone, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("user not exists")
		}
		return User{}, err
	}
	return user, err
}
