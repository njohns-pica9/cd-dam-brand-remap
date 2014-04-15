package main

import (
    _ "github.com/lib/pq"
    "database/sql"
)

type Repository struct {
    DB *sql.DB
}

func (r *Repository) NewConnection(connection string) error {
    db, err := sql.Open("postgres", connection)

    if err != nil {
        return err
    }

    r.DB = db

    return nil
}

func (r *Repository) Close() {
    r.DB.Close()
}

func (r *Repository) RemapBrandDatapoint(mapping map[string]interface{}) {
}