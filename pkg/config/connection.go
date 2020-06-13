package config

import (
	"database/sql"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

	// Driver PostgeSQL
	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
)

// GetConnection ...
func GetConnection() *sql.DB {
	cnf := LoadVariables()
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cnf.DBUser,
			cnf.DBPassword,
			cnf.DBServer,
			cnf.DBPort,
			cnf.DBName,
		)
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Can not connect to database: %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("Database error: %v", err)
		}
	})

	return db
}
