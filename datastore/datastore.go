// Package datastore provides interfaces in which is used to persist the data of DMARC and TLS-RPT reports.
// The implementation should implement:
//  1. For TLS-RPT reports: mailweave.TlsRptMonitoringReports and mailweave.TlsRptMonitoringSources
//  2. For DMARC reports: mailweave.DmarcMonitoringReports and mailweave.DmarcMonitoringSources
//  3. If they require a certain database migration to be executed: Migrator
package datastore

import "context"

// MigrateDirection represents the direction of a database migration, typically used to specify up or down migration.
type MigrateDirection uint8

const (
	// MigrateDirectionUp represents the upward direction for a database migration, indicating the application of changes.
	MigrateDirectionUp MigrateDirection = iota

	// MigrateDirectionDown represents the downward direction for a database migration, indicating the rollback of changes.
	MigrateDirectionDown
)

// Migrator defines an interface for performing database migrations in a specified direction using a given context.
type Migrator interface {
	// Migrate performs database migration in the specified direction (up or down) based on the provided context.
	Migrate(ctx context.Context, direction MigrateDirection) error
}
