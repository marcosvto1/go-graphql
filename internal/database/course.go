package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
	Category    Category
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()
	result, err := c.db.Exec("INSERT INTO course (id, name, description, categoryId) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}

	if _, err := result.RowsAffected(); err != nil {
		return Course{}, err
	}

	c.ID = id

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT cs.id, cs.name, cs.description, c.id, c.name FROM course cs INNER JOIN categories c ON c.id = cs.categoryId")
	if err != nil {
		return nil, err
	}

	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Category.ID, &course.Category.Name)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT cs.id, cs.name, cs.description, c.id, c.name FROM course cs INNER JOIN categories c ON c.id = cs.categoryId WHERE c.id = $1", categoryId)
	if err != nil {
		return nil, err
	}

	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Category.ID, &course.Category.Name)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
