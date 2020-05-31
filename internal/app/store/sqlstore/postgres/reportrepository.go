package postgres

import (
	"errors"
)

//ReportRepository is a reports sql store
type ReportRepository struct {
	store *Store
}

//CreateReport is a function for creating report
func (r *ReportRepository) CreateReport(userID int64, msg string) error {
	if len(msg) == 0 {
		return errors.New("Empty message text")
	}

	if _, err := r.store.db.Exec("INSERT INTO reports (user_id, message) VALUES ($1, $2);", userID, msg); err != nil {
		return err
	}

	return nil
}
