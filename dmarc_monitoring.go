package mailweave

import (
	"context"
	"time"
)

type DmarcDnsRecord struct {
	Policy                string
	SubdomainPolicy       string
	Percentage            float64
	AggregateReportingURI []string
	FailureReportingURI   []string
	SPFAlignmentMode      string
	DKIMAlignmentMode     string
}

type DmarcReportRow struct {
	EmailCount       int64
	SourceIP         string
	ResolvedHostname string

	EnvelopeTo   string
	EnvelopeFrom string
	HeaderFrom   string

	SPFDomain string
	SPFResult string
	SPFScope  string

	DKIMDomain   string
	DKIMSelector string
	DKIMResult   string

	DMARCSPFAligned      bool
	DMARCDKIMAligned     bool
	DMARCInferredAligned bool
	DMARCDisposition     string
}

type DmarcReport struct {
	// Report metadata
	DomainOwner      string
	OrganizationName string
	DomainName       string
	ExtraContactInfo string
	ReportId         string
	RangeStart       time.Time
	RangeEnd         time.Time

	// About the report
	ReceivedAt     time.Time
	EmailSender    string
	EmailSubject   string
	ReportFileName string

	// About the content
	TotalNumberOfEmails int64
	Content             string

	// Report rows
	Rows []DmarcReportRow
}

type DmarcSource struct {
	IPAddress                string
	ReportedEmails           int64
	SPFAlignmentPercentage   float64
	DKIMAlignmentPercentage  float64
	DMARCAlignmentPercentage float64
}

type DmarcSources struct {
	DomainOwner      string
	OrganizationName string
	Domain           string
	Sources          []DmarcSource
}

type DmarcMonitoringReports interface {
	GetDmarcReports(ctx context.Context, domain string) ([]DmarcReport, error)
	GetDmarcReportById(ctx context.Context, domain string, reportId string) (DmarcReport, error)
	WriteDmarcReport(ctx context.Context, domain string, report DmarcReport) error
}

type DmarcMonitoringSources interface {
	GetDmarcSources(ctx context.Context, domain string) ([]DmarcSources, error)
	WriteDmarcSourcesAggregate(ctx context.Context, domain string, reports []DmarcReport) error
}
