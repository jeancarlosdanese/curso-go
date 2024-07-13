// internal/database/course.go

package database

import (
	"database/sql"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	query := "INSERT INTO courses (name, description, category_id) VALUES ($1, $2, $3) RETURNING *"
	err := c.db.QueryRow(query, name, description, categoryID).Scan(&c.ID, &c.Name, &c.Description, &c.CategoryId)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	couses := []Course{}
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		couses = append(couses, course)
	}

	return couses, nil
}

func (c *Course) FindByID(id string) (*Course, error) {
	query := "SELECT id, name, description, category_id FROM courses WHERE id = $1"
	err := c.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description, &c.CategoryId)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Course) Update(id, name, description, categoryID string) (*Course, error) {
	query := "UPDATE courses SET name = $1, description = $2, category_id = $3 WHERE id = $4 RETURNING *"
	err := c.db.QueryRow(query, name, description, categoryID, id).Scan(&c.ID, &c.Name, &c.Description, &c.CategoryId)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Course) Delete(id string) error {
	_, err := c.db.Exec("DELETE FROM courses WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cousers := []Course{}
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		cousers = append(cousers, course)
	}

	return cousers, nil
}
