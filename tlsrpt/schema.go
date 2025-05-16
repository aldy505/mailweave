package tlsrpt

import "time"

type DateRange struct {
	StartDateTime time.Time `json:"start-datetime"`
	EndDateTime   time.Time `json:"end-datetime"`
}

type Policy struct {
	// The type of policy that was applied by the sending
	// domain.  Presently, the only three valid choices are "tlsa",
	// "sts", and the literal string "no-policy-found".  It is provided
	// as a string.
	PolicyType string `json:"policy-type"`
	// An encoding of the applied policy as a JSON array
	// of strings, whether it's a TLSA record ([RFC6698], Section 2.3) or
	// an MTA-STS Policy.  Examples follow in the next section.
	PolicyString []string `json:"policy-string"`
	// The Policy Domain against which the MTA-STS or DANE
	// policy is defined.  In the case of Internationalized Domain Names
	// [RFC5891], the domain MUST consist of the Punycode-encoded
	// A-labels [RFC3492] and not the U-labels.
	PolicyDomain string `json:"policy-domain"`
	// In the case where "policy-type" is "sts", it's
	// the pattern of MX hostnames from the applied policy.  It is
	// provided as a JSON array of strings and is interpreted in the same
	// manner as the rules in "MX Host Validation"; see Section 4.1 of
	// [RFC8461].  In the case of Internationalized Domain Names
	// [RFC5891], the domain MUST consist of the Punycode-encoded
	// A-labels [RFC3492] and not the U-labels.
	MxHost []string `json:"mx-host"`
}

type Summary struct {
	// The aggregate count (an integer,
	// encoded as a JSON number) of successfully negotiated TLS-enabled
	// connections to the receiving site.
	TotalSuccessfulSessionCount int64 `json:"total-successful-session-count"`
	// The aggregate count (an integer,
	// encoded as a JSON number) of failures to negotiate a TLS-enabled
	// connection to the receiving site.
	TotalFailureSessionCount int64 `json:"total-failure-session-count"`
}

type FailureDetail struct {
	ResultType   string `json:"result-type"`
	SendingMTAIP string `json:"sending-mta-ip"`
	// The hostname of the receiving MTA MX
	// record with which the Sending MTA attempted to negotiate a
	// STARTTLS connection.
	ReceivingMxHostname string `json:"receiving-mx-hostname"`
	// The HELLO (HELO) or Extended HELLO
	// (EHLO) string from the banner announced during the reported
	// session.
	ReceivingMxHelo string `json:"receiving-mx-helo"`
	// The destination IP address that was used when
	// creating the outbound session.  It is provided as a string
	// representation of an IPv4 (see below) or IPv6 [RFC5952] address in
	// dot-decimal or colon-hexadecimal notation.
	ReceivingIP string `json:"receiving-ip"`
	// The number of (attempted) sessions that
	// match the relevant "result-type" for this section (an integer,
	// encoded as a JSON number).
	FailedSessionCount int64 `json:"failed-session-count"`
	// A URI [RFC3986] that points to
	// additional information around the relevant "result-type".  For
	// example, this URI might host the complete certificate chain
	// presented during an attempted STARTTLS session.
	AdditionalInformation string `json:"additional-information"`
	// A text field to include a TLS-related error
	// code or error message.
	FailureReasonCode string `json:"failure-reason-code"`
}

type TLSPolicy struct {
	Policy         Policy          `json:"policy"`
	Summary        Summary         `json:"summary"`
	FailureDetails []FailureDetail `json:"failure-details"`
}

type Report struct {
	// OrganizationName is the name of the organization responsible for
	// the report.  It is provided as a string.
	OrganizationName string `json:"organization-name"`
	// The date-time indicates the start and end times for
	// the report range.  It is provided as a string formatted according
	// to "Internet Date/Time Format", Section 5.6 of [RFC3339].  The
	// report should be for a full UTC day, 00:00-24:00.
	DateRange DateRange `json:"date-range"`
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
	Policies []TLSPolicy `json:"policies"`
}
