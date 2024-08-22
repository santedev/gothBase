package store

import (
	"database/sql"
	"fmt"
	"project/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

var DB *Storage

func Sqlconfig() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBName,
		config.Envs.DBsslMode)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

//implement methods using storage database
func (s *Storage) GetExample(uid string) ([]string, error) {
	rows, err := s.db.Query("SELECT column FROM table WHERE column = $1", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rowSlice []string
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			return nil, err
		}
		rowSlice = append(rowSlice, str)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rowSlice, nil
}