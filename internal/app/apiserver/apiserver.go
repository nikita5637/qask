package apiserver

import (
	"database/sql"
	"net/http"
	"qask/internal/app/questions/www"
	"qask/internal/app/store/sqlstore/mysql"
	"qask/internal/app/store/sqlstore/postgres"

	"github.com/sirupsen/logrus"
)

// Start ...
func Start(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	db, err := newDB(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		return err
	}

	var server *server
	questions := www.New()
	switch config.DatabaseDriver {
	case "mysql":
		store := mysql.New(db)
		server = newServer(store, questions)
	case "postgres":
		store := postgres.New(db)
		server = newServer(store, questions)
	}

	server.logger.SetLevel(level)

	server.logger.Infof("Server started with params")
	server.logger.Infof("Bind address \"%s\"", config.BindAddr)
	server.logger.Infof("Database Driver \"%s\"", config.DatabaseDriver)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseDriver string, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(databaseDriver, databaseURL)
	if err != nil {
		return nil, err
	}

	//defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
