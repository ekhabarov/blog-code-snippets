package repo

import "database/sql"

type Entity struct {
	ID   int
	Name string
}

type Database interface {
	GetEntity(id int) (*Entity, error)
}

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{db: db}
}

func (DB) GetEntity(id int) (*Entity, error) {
	return &Entity{}, nil
}
