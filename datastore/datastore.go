package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/aldy505/mailweave"
)

// FakeDatastore implements mailweave.TlsRptMonitoringReports,
// mailweave.TlsRptMonitoringSources, mailweave.DmarcMonitoringReports,
// and mailweave.DmarcMonitoringSources. It should be used for testing purposes.
type FakeDatastore struct {
	TlsRptReports []mailweave.TlsRptReport
	TlsRptSources []mailweave.TlsRptSources
	DmarcReports  []mailweave.DmarcReport
	DmarcSources  []mailweave.DmarcSources
}

var _ mailweave.TlsRptMonitoringReports = (*FakeDatastore)(nil)
var _ mailweave.TlsRptMonitoringSources = (*FakeDatastore)(nil)
var _ mailweave.DmarcMonitoringReports = (*FakeDatastore)(nil)
var _ mailweave.DmarcMonitoringSources = (*FakeDatastore)(nil)

// GetDmarcSources implements mailweave.DmarcMonitoringSources.
func (f *FakeDatastore) GetDmarcSources(ctx context.Context, domain string) ([]mailweave.DmarcSources, error) {
	initialSources := []mailweave.DmarcSources{
		{
			DomainOwner:      domain,
			OrganizationName: "",
			Domain:           domain,
			Sources: []mailweave.DmarcSource{
				{
					IPAddress:                "192.0.2.1",
					ReportedEmails:           28849,
					SPFAlignmentPercentage:   97.8,
					DKIMAlignmentPercentage:  97.7,
					DMARCAlignmentPercentage: 98,
				},
				{
					IPAddress:                "192.0.2.2",
					ReportedEmails:           28009,
					SPFAlignmentPercentage:   100,
					DKIMAlignmentPercentage:  100,
					DMARCAlignmentPercentage: 100,
				},
				{
					IPAddress:                "192.0.2.3",
					ReportedEmails:           19463,
					SPFAlignmentPercentage:   100,
					DKIMAlignmentPercentage:  100,
					DMARCAlignmentPercentage: 100,
				},
				{
					IPAddress:                "192.0.2.4",
					ReportedEmails:           5716,
					SPFAlignmentPercentage:   87.6,
					DKIMAlignmentPercentage:  81.4,
					DMARCAlignmentPercentage: 93,
				},
			},
		},
		{
			DomainOwner:      domain,
			OrganizationName: "Sendgrid",
			Domain:           "sendgrid.net",
			Sources: []mailweave.DmarcSource{
				{
					IPAddress:                "149.72.61.220",
					ReportedEmails:           830,
					SPFAlignmentPercentage:   99.9,
					DKIMAlignmentPercentage:  75.4,
					DMARCAlignmentPercentage: 100,
				},
				{
					IPAddress:                "149.72.114.92",
					ReportedEmails:           329,
					SPFAlignmentPercentage:   100,
					DKIMAlignmentPercentage:  77.5,
					DMARCAlignmentPercentage: 100,
				},
			},
		},
		{
			DomainOwner:      domain,
			OrganizationName: "Amazon SES",
			Domain:           "amazonses.com",
			Sources: []mailweave.DmarcSource{
				{
					IPAddress:                "23.251.232.1",
					ReportedEmails:           1642,
					SPFAlignmentPercentage:   0,
					DKIMAlignmentPercentage:  89.4,
					DMARCAlignmentPercentage: 89.4,
				},
				{
					IPAddress:                "23.251.232.2",
					ReportedEmails:           1603,
					SPFAlignmentPercentage:   0,
					DKIMAlignmentPercentage:  89.1,
					DMARCAlignmentPercentage: 89.1,
				},
			},
		},
	}

	if len(f.DmarcSources) > 0 {
		initialSources = append(initialSources, f.DmarcSources...)
	}

	return initialSources, nil
}

// WriteDmarcSourcesAggregate implements mailweave.DmarcMonitoringSources.
func (f *FakeDatastore) WriteDmarcSourcesAggregate(ctx context.Context, domain string, reports []mailweave.DmarcReport) error {
	if len(f.DmarcSources) == 0 {
		f.DmarcSources = make([]mailweave.DmarcSources, 0)
	}

	// the key is domain
	reportMap := make(map[string][]mailweave.DmarcReport)

	// Classify per domain
	for _, report := range reports {
		if report.DomainOwner != domain {
			continue
		}

		if r, ok := reportMap[report.DomainName]; ok {
			reportMap[report.DomainName] = append(r, report)
		} else {
			reportMap[report.DomainName] = []mailweave.DmarcReport{report}
		}
	}

	// the key is domain
	sourcesMap := make(map[string][]mailweave.DmarcSource)
	organizationNameMap := make(map[string]string)

	for domain, reports := range reportMap {
		// Aggregate per IP address
		type aggregate struct {
			count               int64
			totalReportedEmails int64
			totalSPFAligned     int64
			totalDKIMAligned    int64
			totalDMARCAligned   int64
		}
		// the key is IP address
		ipAggregates := make(map[string]aggregate)
		for _, report := range reports {
			organizationNameMap[domain] = report.OrganizationName

			for _, source := range report.Rows {
				if ipAggregate, ok := ipAggregates[source.SourceIP]; ok {
					ipAggregate.count += 1
					ipAggregate.totalReportedEmails += source.EmailCount
					if source.DMARCInferredAligned {
						ipAggregate.totalDMARCAligned += 1
					}

					if source.SPFResult == "pass" {
						ipAggregate.totalSPFAligned += 1
					}

					if source.DKIMResult == "pass" {
						ipAggregate.totalDKIMAligned += 1
					}

					ipAggregates[source.SourceIP] = ipAggregate
				} else {
					ipAggregate := aggregate{
						count:               1,
						totalReportedEmails: source.EmailCount,
						totalSPFAligned:     0,
						totalDKIMAligned:    0,
						totalDMARCAligned:   0,
					}

					if source.DMARCInferredAligned {
						ipAggregate.totalDMARCAligned += 1
					}

					if source.SPFResult == "pass" {
						ipAggregate.totalSPFAligned += 1
					}

					if source.DKIMResult == "pass" {
						ipAggregate.totalDKIMAligned += 1
					}

					ipAggregates[source.SourceIP] = ipAggregate
				}
			}
		}

		for ipAddress, aggregate := range ipAggregates {
			source := mailweave.DmarcSource{
				IPAddress:                ipAddress,
				ReportedEmails:           aggregate.totalReportedEmails,
				SPFAlignmentPercentage:   float64(aggregate.totalSPFAligned) / float64(aggregate.count) * 100,
				DKIMAlignmentPercentage:  float64(aggregate.totalDKIMAligned) / float64(aggregate.count) * 100,
				DMARCAlignmentPercentage: float64(aggregate.totalDMARCAligned) / float64(aggregate.count) * 100,
			}

			if _, ok := sourcesMap[domain]; !ok {
				sourcesMap[domain] = []mailweave.DmarcSource{source}
			} else {
				sourcesMap[domain] = append(sourcesMap[domain], source)
			}
		}
	}

	for domainName, sources := range sourcesMap {
		// If the domain already exists, we replace it
		for i, source := range f.DmarcSources {
			if source.Domain == domainName {
				f.DmarcSources[i] = mailweave.DmarcSources{
					OrganizationName: organizationNameMap[source.OrganizationName],
					Domain:           domainName,
					Sources:          sources,
					DomainOwner:      domain,
				}
				break
			}
		}

		// If the domain doesn't exist, we add it
		f.DmarcSources = append(f.DmarcSources, mailweave.DmarcSources{
			OrganizationName: organizationNameMap[domainName],
			Domain:           domainName,
			Sources:          sources,
			DomainOwner:      domain,
		})
	}

	return nil
}

// GetDmarcReportById implements mailweave.DmarcMonitoringReports.
func (f *FakeDatastore) GetDmarcReportById(ctx context.Context, domain string, reportId string) (mailweave.DmarcReport, error) {
	initialReports := []mailweave.DmarcReport{
		{
			DomainOwner:         domain,
			OrganizationName:    "Google",
			DomainName:          "google.com",
			ExtraContactInfo:    "noreply-dmarc-support@google.com",
			ReportId:            "8639335954371369510",
			RangeStart:          time.Date(2025, time.May, 13, 0, 0, 0, 0, time.UTC),
			RangeEnd:            time.Date(2025, time.May, 13, 23, 59, 59, 0, time.UTC),
			ReceivedAt:          time.Date(2025, time.May, 14, 0, 0, 0, 0, time.UTC),
			EmailSender:         "dmarc-report@google.com",
			EmailSubject:        "DMARC report for " + domain,
			ReportFileName:      "google.com!example.com!1747008000!1747094399",
			TotalNumberOfEmails: 1000,
			Content:             "",
			Rows:                []mailweave.DmarcReportRow{},
		},
	}

	for _, report := range initialReports {
		if report.DomainOwner == domain && report.ReportId == reportId {
			return report, nil
		}
	}

	for _, report := range f.DmarcReports {
		if report.DomainOwner == domain && report.ReportId == reportId {
			return report, nil
		}
	}

	return mailweave.DmarcReport{}, fmt.Errorf("report not found")
}

// GetDmarcReports implements mailweave.DmarcMonitoringReports.
func (f *FakeDatastore) GetDmarcReports(ctx context.Context, domain string) ([]mailweave.DmarcReport, error) {
	var reports []mailweave.DmarcReport

	for _, report := range f.DmarcReports {
		if report.DomainOwner == domain {
			reports = append(reports, report)
		}
	}

	return reports, nil
}

// WriteDmarcReport implements mailweave.DmarcMonitoringReports.
func (f *FakeDatastore) WriteDmarcReport(ctx context.Context, domain string, report mailweave.DmarcReport) error {
	if len(f.DmarcReports) == 0 {
		f.DmarcReports = make([]mailweave.DmarcReport, 0)
	}
	report.DomainOwner = domain
	f.DmarcReports = append(f.DmarcReports, report)
	return nil
}

// GetTlsRptSources implements mailweave.TlsRptMonitoringSources.
func (f *FakeDatastore) GetTlsRptSources(ctx context.Context, domain string) ([]mailweave.TlsRptSources, error) {
	var sources []mailweave.TlsRptSources

	for _, source := range f.TlsRptSources {
		if source.DomainOwner == domain {
			sources = append(sources, source)
		}
	}

	return sources, nil
}

// WriteTlsRptSourcesAggregate implements mailweave.TlsRptMonitoringSources.
func (f *FakeDatastore) WriteTlsRptSourcesAggregate(ctx context.Context, domain string, reports []mailweave.TlsRptReport) error {
	if len(f.TlsRptSources) == 0 {
		f.TlsRptSources = make([]mailweave.TlsRptSources, 0)
	}

	// the key is domain
	reportMap := make(map[string][]mailweave.TlsRptReport)

	// Classify per domain
	for _, report := range reports {
		if report.DomainOwner != domain {
			continue
		}

		if r, ok := reportMap[report.DomainName]; ok {
			reportMap[report.DomainName] = append(r, report)
		} else {
			reportMap[report.DomainName] = []mailweave.TlsRptReport{report}
		}
	}

	// the key is domain
	sourcesMap := make(map[string]mailweave.TlsRptSources)
	organizationNameMap := make(map[string]string)

	for domainName, reports := range reportMap {
		// Aggregate per IP address
		type aggregate struct {
			totalSuccessful int64
			totalFailed     int64
		}
		// the key is IP address
		ipAggregates := make(map[string]aggregate)
		for _, report := range reports {
			organizationNameMap[domainName] = report.OrganizationName

			for _, source := range report.Rows {
				if ipAggregate, ok := ipAggregates[source.IPAddress]; ok {
					ipAggregate.totalSuccessful += source.SuccessfulSessionCount
					ipAggregate.totalFailed += source.FailedSessionCount
					ipAggregates[source.IPAddress] = ipAggregate
				} else {
					ipAggregate := aggregate{
						totalSuccessful: source.SuccessfulSessionCount,
						totalFailed:     source.FailedSessionCount,
					}
					ipAggregates[source.IPAddress] = ipAggregate
				}
			}
		}

		for _, aggregate := range ipAggregates {
			source := mailweave.TlsRptSources{
				DomainOwner:                 domain,
				OrganizationName:            organizationNameMap[domainName],
				Domain:                      domainName,
				SuccessfulSessionPercentage: float64(aggregate.totalSuccessful) / float64(aggregate.totalSuccessful+aggregate.totalFailed) * 100,
			}

			sourcesMap[domainName] = source
		}
	}

	for domainName, sources := range sourcesMap {
		// If the domain already exists, we replace it
		for i, source := range f.TlsRptSources {
			if source.Domain == domainName {
				f.TlsRptSources[i] = sources
				break
			}
		}

		// If the domain doesn't exist, we add it
		f.TlsRptSources = append(f.TlsRptSources, sources)
	}

	return nil
}

// GetTlsRptReportById implements mailweave.TlsRptMonitoringReports.
func (f *FakeDatastore) GetTlsRptReportById(ctx context.Context, domain string, reportId string) (mailweave.TlsRptReport, error) {
	for _, report := range f.TlsRptReports {
		if report.DomainOwner == domain && report.ReportId == reportId {
			return report, nil
		}
	}

	return mailweave.TlsRptReport{}, fmt.Errorf("report not found")
}

// GetTlsRptReports implements mailweave.TlsRptMonitoringReports.
func (f *FakeDatastore) GetTlsRptReports(ctx context.Context, domain string) ([]mailweave.TlsRptReport, error) {
	var reports []mailweave.TlsRptReport

	for _, report := range f.TlsRptReports {
		if report.DomainOwner == domain {
			reports = append(reports, report)
		}
	}

	return reports, nil
}

// WriteTlsRptReport implements mailweave.TlsRptMonitoringReports.
func (f *FakeDatastore) WriteTlsRptReport(ctx context.Context, domain string, report mailweave.TlsRptReport) error {
	if len(f.TlsRptReports) == 0 {
		f.TlsRptReports = make([]mailweave.TlsRptReport, 0)
	}
	report.DomainOwner = domain
	f.TlsRptReports = append(f.TlsRptReports, report)
	return nil
}
