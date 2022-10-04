package users

type User struct {
	Id     int
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

var users []User
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, email string, age int, height int, active bool, date string) (User, error)
	LastID() (int, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name string, email string, age int, height int, active bool, date string) (User, error) {
	user := User{id, name, email, age, height, active, date}
	users = append(users, user)
	lastID = user.Id
	return user, nil
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
