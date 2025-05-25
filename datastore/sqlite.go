package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aldy505/mailweave"
)

// SqliteDatastore represents a datastore implemented with SQLite for managing and querying email monitoring data.
type SqliteDatastore struct {
	db *sql.DB
}

var _ Migrator = (*SqliteDatastore)(nil)
var _ mailweave.TlsRptMonitoringReports = (*SqliteDatastore)(nil)
var _ mailweave.TlsRptMonitoringSources = (*SqliteDatastore)(nil)
var _ mailweave.DmarcMonitoringReports = (*SqliteDatastore)(nil)
var _ mailweave.DmarcMonitoringSources = (*SqliteDatastore)(nil)

// NewSqliteDatastore initializes a new SqliteDatastore with the provided *sql.DB connection.
// Returns an error if the provided database connection is nil.
func NewSqliteDatastore(db *sql.DB) (*SqliteDatastore, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return &SqliteDatastore{
		db: db,
	}, nil
}

func (s *SqliteDatastore) GetDmarcSources(ctx context.Context, domain string) ([]mailweave.DmarcSources, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) WriteDmarcSourcesAggregate(ctx context.Context, domain string, reports []mailweave.DmarcReport) error {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) GetDmarcReports(ctx context.Context, domain string) ([]mailweave.DmarcReport, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) GetDmarcReportById(ctx context.Context, domain string, reportId string) (mailweave.DmarcReport, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) WriteDmarcReport(ctx context.Context, domain string, report mailweave.DmarcReport) error {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) GetTlsRptSources(ctx context.Context, domain string) ([]mailweave.TlsRptSources, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) WriteTlsRptSourcesAggregate(ctx context.Context, domain string, reports []mailweave.TlsRptReport) error {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) GetTlsRptReports(ctx context.Context, domain string) ([]mailweave.TlsRptReport, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) GetTlsRptReportById(ctx context.Context, domain string, reportId string) (mailweave.TlsRptReport, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) WriteTlsRptReport(ctx context.Context, domain string, report mailweave.TlsRptReport) error {
	// TODO implement me
	panic("implement me")
}

func (s *SqliteDatastore) Migrate(ctx context.Context, direction MigrateDirection) error {
	// TODO implement me
	panic("implement me")
}
