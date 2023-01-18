package repository

import (
	"database/sql"
	"final-project-sanbercode/structs"
)

// UserRepo is the repository for user
type UserRepo struct {
	DB *sql.DB
}

// NewUserRepo is the constructor for user repository
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// GetAll is the function to get all users
func (u *UserRepo) GetAll() ([]structs.Users, error) {
	var users []structs.Users

	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user structs.Users
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Remember_token, &user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID is the function to get user by id
func (u *UserRepo) GetByID(id int) (structs.Users, error) {
	var user structs.Users

	err := u.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Login
func (u *UserRepo) Login(email string, password string) (structs.Users, error) {
	var user structs.Users

	err := u.DB.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil

}

// Insert is the function to insert user
func (u *UserRepo) Insert(user structs.Users) (structs.Users, error) {
	err := u.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Update is the function to update user with id
func (u *UserRepo) Update(id int, user structs.Users) (structs.Users, error) {
	err := u.DB.QueryRow("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id", user.Name, user.Email, user.Password, id).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete is the function to delete user
func (u *UserRepo) Delete(id int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
