package teststore

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

	return nil
}
