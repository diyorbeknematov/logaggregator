package sqlite

import (
	"database/sql"
	"fmt"
	"logaggregator/models"
)

type LogRepository interface {
	SaveLog(log models.Log) error
	GetLogs(fLogs models.Filter) ([]models.Log, error)
	DeleteLogs(flogs models.Filter) error
}

type logRepositoryImpl struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) LogRepository {
	return &logRepositoryImpl{
		db: db,
	}
}

func (repo *logRepositoryImpl) SaveLog(log models.Log) error {
	query := `
		INSERT INTO logs (
			timestamp,
			user_id,
			level,
			message,
			service,
			error
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := repo.db.Exec(query, log.Timestamp, log.UserID, log.Level, log.Message, log.Service, log.Error)
	if err != nil {
		return err
	}

	return nil
}

func (repo *logRepositoryImpl) GetLogs(fLogs models.Filter) ([]models.Log, error) {
	var (
		args   []interface{}
		filter string
	)

	query := `
		SELECT 
			timestamp,
			user_id,
			level,
			message,
			service,
			error
		FROM
			logs
		WHERE
			1=1
	`

	if fLogs.TimestampFrom != "" && fLogs.TimestampTo != "" {
		filter += " AND timestamp BETWEEN ? AND ?"
		args = append(args, fLogs.TimestampFrom, fLogs.TimestampTo)
	}

	if fLogs.Level != "" {
		filter += " AND level LIKE ?"
		args = append(args, "%"+fLogs.Level+"%")
	}

	if fLogs.Service != "" {
		filter += " AND service LIKE ?"
		args = append(args, "%"+fLogs.Service+"%")
	}

	if fLogs.UserID != "" {
		filter += " AND user_id LIKE ?"
		args = append(args, "%"+fLogs.UserID+"%")
	}

	query += filter
	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var log models.Log

		err := rows.Scan(&log.Timestamp, &log.UserID, &log.Level, &log.Message, &log.Service, &log.Error)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (repo *logRepositoryImpl) DeleteLogs(flogs models.Filter) error {
	var (
		query string
		args  []interface{}
	)

	switch {
	case flogs.TimestampFrom != "" && flogs.TimestampTo != "":
		query = "DELETE FROM logs WHERE timestamp BETWEEN ? AND ?;"
		args = append(args, flogs.TimestampFrom, flogs.TimestampTo)
	case flogs.UserID != "":
		query = "DELETE FROM logs WHERE user_id = ?;"
		args = append(args, flogs.UserID)
	case flogs.Service != "":
		query = "DELETE FROM logs WHERE service = ?;"
		args = append(args, flogs.Service)
	case flogs.Level != "":
		query = "DELETE FROM logs WHERE level = ?;"
		args = append(args, flogs.Level)
	default:
		fmt.Println("No filter specified, not performing delete operation.")
		return fmt.Errorf("No filter specified, not performing delete operation.")
	}

	// Execute DELETE query
	result, err := repo.db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return err
	}

	fmt.Printf("Deleted %d rows\n", rowsAffected)

	return nil
}
