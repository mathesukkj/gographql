package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type CategoryDb struct {
	db *sql.DB
}

type Category struct {
	ID          string
	Name        string
	Description string
}

func NewCategoryDb(db *sql.DB) *CategoryDb {
	return &CategoryDb{db: db}
}

func (c *CategoryDb) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id,
		name,
		description,
	)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
