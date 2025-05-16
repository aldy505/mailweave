package mailweave

import (
	"context"
	"time"
)

type TlsRptDnsRecord struct {
	ReportsMailTransport  []string
	ReportsHttpsTransport []string
}

type TlsRptReportRow struct {
	DomainName   string
	IPAddress    string
	PolicyType   string
	PolicyString []string
	MxHost       []string

	SuccessfulSessionCount int64
	FailedSessionCount     int64
}

type TlsRptReport struct {
	// Metadata
	DomainOwner      string
	OrganizationName string
	DomainName       string
	ReportId         string
	ExtraContactInfo string
	RangeStart       time.Time
	RangeEnd         time.Time

	// About the report
	ReceivedAt     time.Time
	EmailSender    string
	EmailSubject   string
	ReportFileName string

	// Content
	TotalNumberOfSessions int64
	Content               string

	// Report rows
	Rows []TlsRptReportRow
}

type TlsRptSources struct {
	DomainOwner                 string
	OrganizationName            string
	Domain                      string
	SuccessfulSessionPercentage float64
}

type TlsRptMonitoringReports interface {
	GetTlsRptReports(ctx context.Context, domain string) ([]TlsRptReport, error)
	GetTlsRptReportById(ctx context.Context, domain string, reportId string) (TlsRptReport, error)
	WriteTlsRptReport(ctx context.Context, domain string, report TlsRptReport) error
}

type TlsRptMonitoringSources interface {
	GetTlsRptSources(ctx context.Context, domain string) ([]TlsRptSources, error)
	WriteTlsRptSourcesAggregate(ctx context.Context, domain string, reports []TlsRptReport) error
}
