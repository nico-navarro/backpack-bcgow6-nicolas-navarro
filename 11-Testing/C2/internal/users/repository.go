package users

import (
	"C2/pkg/store"
	"errors"
)

type User struct {
	Id     int
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

// ***Importado**//
type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, email string, age int, height int, active bool, date string) (User, error)
	LastID() (int, error)
	Update(id int, name string, email string, age int, height int, active bool, date string) (User, error)
	Delete(id int) (User, error)
	Patch(id int, name string, email string, age *int, height *int, active *bool, date string) (User, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, name string, email string, age int, height int, active bool, date string) (User, error) {
	var users []User
	r.db.Read(&users)

	user := User{id, name, email, age, height, active, date}
	users = append(users, user)
	err := r.db.Write(users)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	r.db.Read(&users)
	return users, nil
}

func (r *repository) LastID() (int, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, nil
	}
	return users[len(users)-1].Id, nil
}

func (r *repository) Update(id int, name string, email string, age int, height int, active bool, date string) (User, error) {
	// Load users
	var users []User
	r.db.Read(&users)

	updatedUser := User{}
	for i, user := range users {
		if user.Id == id {
			updatedUser = User{id, name, email, age, height, active, date}
			users[i] = updatedUser
			//Save users
			err := r.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return updatedUser, nil
		}
	}
	return updatedUser, errors.New("no existe el ID especificado")
}

func (r *repository) Delete(id int) (User, error) {
	// Load users
	var users []User
	r.db.Read(&users)

	deletedUser := User{}
	for i, user := range users {
		if user.Id == id {
			deletedUser = users[i]
			users = append(users[:i], users[i+1:]...)
			//Save users
			err := r.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return deletedUser, nil
		}
	}
	return deletedUser, errors.New("no existe el ID especificado")
}

func (r *repository) Patch(id int, name string, email string, age *int, height *int, active *bool, date string) (User, error) {
	// Load users
	var users []User
	r.db.Read(&users)

	for i, user := range users {
		if user.Id == id {
			if name != "" {
				users[i].Name = name
			}
			if email != "" {
				users[i].Email = email
			}
			if age != nil {
				users[i].Age = *age
			}
			if height != nil {
				users[i].Height = *height
			}
			if active != nil {
				users[i].Active = *active
			}
			if date != "" {
				users[i].Date = date
			}
			//Save users
			err := r.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return users[i], nil
		}
	}
	return User{}, errors.New("no existe el ID especificado")
}
