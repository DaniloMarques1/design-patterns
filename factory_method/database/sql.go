package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqlDatabase struct {
	db *sql.DB
}

const CREATE_TABLES = `
create table if not exists person(
	id string primary key,
	name varchar(100) not null unique
);
`

func NewSqlDatabase() (*SqlDatabase, error) {
	db, err := sql.Open("sqlite3", "db.sql")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(CREATE_TABLES); err != nil {
		return nil, err
	}

	return &SqlDatabase{db: db}, nil
}

func (s *SqlDatabase) Save(person *Person) error {
	stmt, err := s.db.Prepare("insert into person(id, name) values($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(person.Id, person.Name); err != nil {
		return err
	}

	return nil
}

func (s *SqlDatabase) FindByName(name string) (*Person, error) {
	stmt, err := s.db.Prepare("select id, name from person where name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	person := &Person{}
	if err := stmt.QueryRow(name).Scan(&person.Id, &person.Name); err != nil {
		return nil, err
	}

	return person, nil
}
