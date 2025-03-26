package postgres

import (
	"database/sql"
	"fmt"
)

type Role struct {
	db *sql.DB
}

func NewRole(db *sql.DB) *Role {
	return &Role{db: db}
}

func (r *Role) GetLevelAccess(roleID int) (int, error) {
	query := fmt.Sprintf("SELECT level_access FROM %s WHERE role_id = $1", rolesTable)
	row := r.db.QueryRow(query, roleID)
	var levelAccess int
	err := row.Scan(&levelAccess)
	if err != nil {
		return -1, fmt.Errorf("failed to retrieve level access from role: %w", err)
	}
	return levelAccess, nil
}
