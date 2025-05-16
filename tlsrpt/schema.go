package tlsrpt

type dateRange struct {
	StartDateTime string `json:"start-datetime"`
	EndDateTime   string `json:"end-datetime"`
}

type Policy struct {
	PolicyType   string   `json:"policy-type"`
	PolicyString []string `json:"policy-string"`
	PolicyDomain string   `json:"policy-domain"`
	MxHost       []string `json:"mx-host"`
}

type Summary struct {
	TotalSuccessfulSessionCount int `json:"total-successful-session-count"`
	TotalFailureSessionCount    int `json:"total-failure-session-count"`
}

type FailureDetail struct {
	ResultType            string `json:"result-type"`
	SendingMTAIP          string `json:"sending-mta-ip"`
	ReceivingMxHostname   string `json:"receiving-mx-hostname"`
	ReceivingMxHelo       string `json:"receiving-mx-helo"`
	ReceivingIP           string `json:"receiving-ip"`
	FailedSessionCount    int    `json:"failed-session-count"`
	AdditionalInformation string `json:"additional-information"`
	FailureReasonCode     string `json:"failure-reason-code"`
}

type policy struct {
	Policy         Policy          `json:"policy"`
	Summary        Summary         `json:"summary"`
	FailureDetails []FailureDetail `json:"failure-details"`
}

type schema struct {
	// OrganizationName is the name of the organization responsible for
	// the report.  It is provided as a string.
	OrganizationName string `json:"organization-name"`
	// The date-time indicates the start and end times for
	// the report range.  It is provided as a string formatted according
	// to "Internet Date/Time Format", Section 5.6 of [RFC3339].  The
	// report should be for a full UTC day, 00:00-24:00.
	DateRange dateRange `json:"date-range"`
	// The contact information for the party responsible
	// for the report.  It is provided as a string formatted according to
	// "Addr-Spec Specification", Section 3.4.1 of [RFC5322].
	ContactInfo string `json:"contact-info"`
	// A unique identifier for the report.  Report authors
	// may use whatever scheme they prefer to generate a unique
	// identifier.  It is provided as a string.
	ReportID string `json:"report-id"`
	// The type of policy that was applied by the sending
	//   domain.  Presently, the only three valid choices are "tlsa",
	//   "sts", and the literal string "no-policy-found".  It is provided
	//   as a string.
	Policies []policy `json:"policies"`
}
