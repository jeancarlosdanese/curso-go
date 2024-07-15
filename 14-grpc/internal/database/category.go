// internal/database/category.go

package database

import (
	"database/sql"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (*Category, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING *"
	err := c.db.QueryRow(query, name, description).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByID(id string) (*Category, error) {
	err := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) FindByCourseID(courseID string) (*Category, error) {
	err := c.db.QueryRow("SELECT categories.id, categories.name, categories.description FROM categories JOIN courses ON categories.id = courses.category_id WHERE courses.id = $1", courseID).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}

	return c, nil
}
