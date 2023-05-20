package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	result, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES($1,$2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	if _, err := result.RowsAffected(); err != nil {
		return Category{}, err
	}

	c.ID = id

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}

	var categories []Category
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {
	var id, name, description string
	sql := `
        SELECT c.id, c.name, c.description
        FROM categories c JOIN course co ON c.id = co.categoryId
        WHERE co.id = $1
    `
	err := c.db.QueryRow(sql, courseId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
