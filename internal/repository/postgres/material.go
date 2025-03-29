package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/models"
)

type Material struct {
	db *sql.DB
}

func NewMaterial(db *sql.DB) *Material {
	return &Material{
		db: db,
	}
}

func (m *Material) CreateMaterial(courseID int, title string, content string) error {
	query := fmt.Sprintf("INSERT INTO %s(course_id, title, content) VALUES($1, $2, $3))", materialsTable)
	_, err := m.db.Exec(query, courseID, title, content)
	return err
}

func (m *Material) GetAllMaterialByID(courseID int) ([]models.Material, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE course_id=$1", materialsTable)
	rows, err := m.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var materials []models.Material
	for rows.Next() {
		var material models.Material
		rows.Scan(&material.ID, &material.Title, &material.Content)
		materials = append(materials, material)
	}
	return materials, nil
}

func (m *Material) GetMaterialByID(materialID int) (models.Material, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", materialsTable)
	row := m.db.QueryRow(query, materialID)
	var material models.Material
	err := row.Scan(&material.ID, &material.Title, &material.Content)
	return material, err
}

func (m *Material) UpdateTitleMaterial(materialID int, title string) error {
	query := fmt.Sprintf("UPDATE  %s SET title = $1 WHERE course_id = $2", materialsTable)
	_, err := m.db.Exec(query, title, materialID)
	return err
}

func (m *Material) UpdateContentMaterial(materialID int, content string) error {
	query := fmt.Sprintf("UPDATE %ss SET content = $1 WHERE course_id = $2", materialsTable)
	_, err := m.db.Exec(query, content, materialID)
	return err
}

func (m *Material) DeleteMaterial(materialID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE course_id = $1", materialsTable)
	_, err := m.db.Exec(query, materialID)
	return err
}
