package postgres

import (
	"database/sql"
	"fmt"
	"github.com/BioMihanoid/LearningManagementSystem/internal/models"
)

type Log struct {
	db *sql.DB
}

func NewLog(db *sql.DB) *Log {
	return &Log{
		db: db,
	}
}

func (l *Log) CreateLog(userID int, action string) error {
	query := fmt.Sprintf("INSERT %s(user_id, action) VALUES($1, $2)", logsTable)
	_, err := l.db.Exec(query, userID, action)
	return err
}

func (l *Log) GetLogByID(logID int) (models.Log, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE log_id=$1", logsTable)
	row := l.db.QueryRow(query, logID)
	var log models.Log
	err := row.Scan(&log.ID, &log.UserID, &log.Action, &log.Timestamp)
	return log, err
}

func (l *Log) GetLogsCurrentUser(userID int) ([]models.Log, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", logsTable)
	rows, err := l.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logs []models.Log
	for rows.Next() {
		var log models.Log
		err = rows.Scan(&log.ID, &log.UserID, &log.Action, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (l *Log) DeleteLogByID(logID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE log_id=$1", logsTable)
	_, err := l.db.Exec(query, logID)
	return err
}
