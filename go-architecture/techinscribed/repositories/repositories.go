package repositories

import (
	"database/sql"
	"fmt"
	"techinscribed-course/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindById(id int) (*models.User, error) {
	var user models.User

	row := r.db.QueryRow("select * from employee where id = $1", id)
	// fmt.Println("u", id, user)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		fmt.Println("error naman", err)
		return &models.User{}, nil
	}

	// fmt.Println("here")
	return &user, nil
}

func (r *UserRepo) Save(user *models.User) error {
	return nil
}
