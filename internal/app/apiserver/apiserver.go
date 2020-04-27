package apiserver

import (
	"database/sql"
	"net/http"
	"qask/internal/app/questions/www"
	"qask/internal/app/store/teststore"

	"github.com/sirupsen/logrus"
)

// Start ...
func Start(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	// db, err := newDB(config.DatabaseURL)
	// if err != nil {
	// return err
	// }

	store := teststore.New()
	questions := www.New()

	server := newServer(store, questions)
	server.logger.SetLevel(level)

	server.logger.Infof("Server started with params")
	server.logger.Infof("Bind address \"%s\"", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
