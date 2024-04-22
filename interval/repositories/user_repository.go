package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-jwt/interval/models"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type userRepository struct {
	db *DB
}

func NewUserRepository(db *DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := goqu.Insert("user").
		Rows(goqu.Record{
			"email":   user.Email,
			"name":    user.Name,
			"surname": user.Surname,
		}).
		Returning("id")

	insertSQL, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	err = r.db.Conn.QueryRowContext(ctx, insertSQL, args...).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetById(ctx context.Context, userID int) (*models.User, error) {
	if userID < 1 {
		return nil, errors.New("record not found")
	}

	query := goqu.From("user").
		Where(goqu.Ex{"id": userID})

	insertSql, args, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	user := models.User{}
	err = r.db.Conn.QueryRowContext(ctx, insertSql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Surname,
	)
	if err != nil {
		if errors.Is(err, errors.New("no rows in result set")) {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	if _, err := r.GetById(ctx, int(user.ID)); err != nil {
		return err
	}

	query := goqu.Update("user").
		Set(goqu.Record{
			"email":   user.Email,
			"name":    user.Name,
			"surname": user.Surname,
		}).
		Where(goqu.Ex{"id": user.ID})

	updateSQL, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = r.db.Conn.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, userID int) error {
	if _, err := r.GetById(ctx, int(userID)); err != nil {
		return err
	}

	query := goqu.Delete("user").
		Where(goqu.Ex{"id": userID})

	deleteSQL, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = r.db.Conn.ExecContext(ctx, deleteSQL, args...)
	if err != nil {
		return err
	}

	return nil
}
