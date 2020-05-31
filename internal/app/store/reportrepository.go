package store

//ReportRepository is a interface
//for working with reports
type ReportRepository interface {
	CreateReport(int64, string) error
}
