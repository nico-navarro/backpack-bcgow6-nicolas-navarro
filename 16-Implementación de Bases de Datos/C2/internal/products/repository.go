package products

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Repository interface {
	Store(p domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

const (
	GET_BY_NAME = "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	STORE       = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)"
)

func (r *repository) Store(p domains.Product) (int, error) {
	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}

func (r *repository) GetByName(name string) (domains.Product, error) {
	stmt, err := r.db.Prepare(GET_BY_NAME)
	if err != nil {
		return domains.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var product domains.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domains.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}
