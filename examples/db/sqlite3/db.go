package sqlite3

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/ldej/go-training/examples/db"

	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	db *sql.DB
}

func NewDB(filename string) (db.DB, error) {
	sqliteDB, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	err = sqliteDB.Ping()
	if err != nil {
		return nil, err
	}

	s := &service{
		db: sqliteDB,
	}
	err = s.createTable()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *service) createTable() error {
	statement := `
		CREATE TABLE IF NOT EXISTS things (
			uuid text primary key,
			name text,
			value text,
			created timestamp,
			updated timestamp
		)`
	_, err := s.db.Exec(statement)
	return err
}

func (s *service) GetThing(uuid string) (db.Thing, error) {
	query := `SELECT uuid, name, value, updated, created FROM things WHERE uuid=?`
	row := s.db.QueryRow(query, uuid)

	var thing db.Thing
	err := row.Scan(&thing.UUID, &thing.Name, &thing.Value, &thing.Updated, &thing.Created)
	if err == sql.ErrNoRows {
		return db.Thing{}, db.ErrThingNotFound
	}
	if err != nil {
		return db.Thing{}, err
	}
	return thing, nil
}

func (s *service) CreateThing(name string, value string) (db.Thing, error) {
	now := time.Now().UTC()
	thing := db.Thing{
		UUID:    uuid.New().String(),
		Name:    name,
		Value:   value,
		Updated: now,
		Created: now,
	}
	query := `INSERT INTO things (uuid, name, value, created, updated) 
		VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(query, thing.UUID, thing.Name, thing.Value, thing.Created, thing.Updated)
	if err != nil {
		return db.Thing{}, err
	}
	return thing, nil
}

func (s *service) UpdateThing(uuid string, value string) (db.Thing, error) {
	query := `UPDATE things SET value=? WHERE uuid=?`
	_, err := s.db.Exec(query, value, uuid)
	if err != nil {
		return db.Thing{}, err
	}
	return s.GetThing(uuid)
}

func (s *service) DeleteThing(uuid string) error {
	query := `DELETE FROM things WHERE uuid=?`
	_, err := s.db.Exec(query, uuid)
	return err
}

func (s *service) ListThings(offset int, limit int) ([]db.Thing, int, error) {
	query := `SELECT uuid, name, value, updated, created FROM things LIMIT ?, ?`
	rows, err := s.db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var things []db.Thing
	for rows.Next() {
		var thing db.Thing
		err = rows.Scan(&thing.UUID, &thing.Name, &thing.Value, &thing.Updated, &thing.Created)
		if err != nil {
			return nil, 0, err
		}
		things = append(things, thing)
	}
	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}

	count, err := s.GetSize()
	if err != nil {
		return nil, 0, err
	}
	return things, count, nil
}

func (s *service) GetSize() (int, error) {
	query := `SELECT COUNT(*) FROM things`
	row := s.db.QueryRow(query)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
