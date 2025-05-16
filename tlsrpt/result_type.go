package tlsrpt

import "encoding/json"

type ResultType struct {
	category ResultTypeCategory
	key      string
	detail   string
}

func (r *ResultType) Category() ResultTypeCategory {
	return r.category
}

func (r *ResultType) Key() string {
	return r.key
}

func (r *ResultType) Detail() string {
	return r.detail
}

func (r *ResultType) String() string {
	return r.key
}

func (r ResultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.key)
}

func (r *ResultType) UnmarshalJSON(b []byte) error {
	var key string
	err := json.Unmarshal(b, &key)
	if err != nil {
		return err
	}

	*r = *ParseResultType(key)
	return nil
}

type ResultTypeCategory string

const (
	CategoryNegotiationFailure ResultTypeCategory = "Negotiation Failure"
	CategoryPolicyFailure      ResultTypeCategory = "Policy Failure"
	CategoryGeneralFailure     ResultTypeCategory = "General Failure"
)

var ResultStartTLSNotSupported = &ResultType{
	category: CategoryNegotiationFailure,
	key:      "starttls-not-supported",
	detail:   "This indicates that the recipient MX did not support STARTTLS.",
}

var ResultCertificateHostMismatch = &ResultType{
	category: CategoryPolicyFailure,
	key:      "certificate-host-mismatch",
	detail:   "This indicates that the certificate presented did not adhere to the constraints specified in the MTA-STS or DANE policy, e.g., if the MX hostname does not match any listed in the subject alternative name (SAN).",
}

var ResultCertificateExpired = &ResultType{
	category: CategoryPolicyFailure,
	key:      "certificate-expired",
	detail:   "This indicates that the certificate presented has expired.",
}

var ResultCertificateNotTrusted = &ResultType{
	category: CategoryPolicyFailure,
	key:      "certificate-not-trusted",
	detail:   "This is a label that covers multiple certificate-related failures that include, but are not limited to, errors such as untrusted/unknown certification authorities (CAs), certificate name constraints, certificate chain errors, etc.  When using this declaration, the reporting MTA SHOULD utilize the \"failure-reason-code\" to provide more information to the receiving entity.",
}

var ResultValidationFailure = &ResultType{
	category: CategoryGeneralFailure,
	key:      "validation-failure",
	detail:   "This indicates a general failure for a reason not matching a category above.  When using this declaration, the reporting MTA SHOULD utilize the \"failure-reason-code\" to provide more information to the receiving entity.",
}

var ResultTLSAInvalid = &ResultType{
	category: CategoryPolicyFailure,
	key:      "tlsa-invalid",
	detail:   "This indicates a validation error in the TLSA record associated with a DANE policy.  None of the records in the RRset were found to be valid.",
}

var ResultDNSSECInvalid = &ResultType{
	category: CategoryPolicyFailure,
	key:      "dnssec-invalid",
	detail:   "This indicates that no valid records were returned from the recursive resolver.",
}

var ResultDANERequired = &ResultType{
	category: CategoryPolicyFailure,
	key:      "dane-required",
	detail:   "This indicates that the sending system is configured to require DANE TLSA records for all the MX hosts of the destination domain, but no DNSSEC-validated TLSA records were present for the MX host that is the subject of the report.  Mandatory DANE for SMTP is described in Section 6 of [RFC7672].  Such policies may be created by mutual agreement between two organizations that frequently exchange sensitive content via email.",
}

var ResultSTSPolicyFetchError = &ResultType{
	category: CategoryPolicyFailure,
	key:      "sts-policy-fetch-error",
	detail:   "This indicates a failure to retrieve an MTA-STS policy, for example, because the policy host is unreachable.",
}

var ResultSTSPolicyInvalid = &ResultType{
	category: CategoryPolicyFailure,
	key:      "sts-policy-invalid",
	detail:   "This indicates a validation error for the overall MTA-STS Policy.",
}

var ResultSTSWebPKIInvalid = &ResultType{
	category: CategoryPolicyFailure,
	key:      "sts-webpki-invalid",
	detail:   "This indicates that the MTA-STS Policy could not be authenticated using PKIX validation.",
}

func ParseResultType(key string) *ResultType {
	switch key {
	case "starttls-not-supported":
		return ResultStartTLSNotSupported
	case "certificate-host-mismatch":
		return ResultCertificateHostMismatch
	case "certificate-expired":
		return ResultCertificateExpired
	case "certificate-not-trusted":
		return ResultCertificateNotTrusted
	case "validation-failure":
		return ResultValidationFailure
	case "tlsa-invalid":
		return ResultTLSAInvalid
	case "dnssec-invalid":
		return ResultDNSSECInvalid
	case "dane-required":
		return ResultDANERequired
	case "sts-policy-fetch-error":
		return ResultSTSPolicyFetchError
	case "sts-policy-invalid":
		return ResultSTSPolicyInvalid
	case "sts-webpki-invalid":
		return ResultSTSWebPKIInvalid
	default:
		return &ResultType{
			category: CategoryGeneralFailure,
			key:      key,
			detail:   "",
		}
	}
}
