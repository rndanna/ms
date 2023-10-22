package postgresql

import (
	"database/sql"
	"engine-service/internal/models"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var ErrQueryRows = errors.New("ErrQueryRows")
var ErrScanRows = errors.New("ErrScanRows")

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) GetEngines(ids []int) ([]models.Engine, error) {
	var engines []models.Engine

	rows, err := db.db.Query(`
		SELECT id, name, description 
		FROM engines
		WHERE id = ANY($1::int[])
	`, pq.Array(ids))
	if err != nil {
		return engines, fmt.Errorf("err GetUserEngines Query: %w", ErrScanRows)
	}
	defer rows.Close()

	for rows.Next() {
		var engine models.Engine

		if err = rows.Scan(&engine.ID, &engine.Name, &engine.Description); err != nil {
			return engines, fmt.Errorf("err GetUserEngines Scan: %w", ErrScanRows)
		}
		engines = append(engines, engine)
	}

	return engines, nil
}

func (db *DB) GetEngine(id string) (*models.Engine, error) {
	var engine models.Engine

	if err := db.db.QueryRow(`
		SELECT id, name, description 
		FROM engines
		WHERE id = $1
    `, id).Scan(&engine.ID, &engine.Name, &engine.Description); err != nil {
		if err == sql.ErrNoRows {
			return &engine, err
		}

		return &engine, fmt.Errorf("err GetEngine QueryRow: %w", ErrScanRows)

	}

	return &engine, nil
}
