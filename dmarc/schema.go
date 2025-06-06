package dmarc

// This file is acquired from https://github.com/desdic/godmarcparser/blob/master/dmarc/dmarc.go
//
// The MIT License (MIT)
//
// Copyright (c) 2018 Kim Gert Nielsen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"bytes"
	"encoding/xml"
	"regexp"
	"time"
)

var brokenschema = regexp.MustCompile(`\<xs:schema[^>]*>`)
var matchschema = regexp.MustCompile(`\<\/xs:schema[^>]*>`)

// Content is the structure for processing data
type Content struct {
	From string
	Name string
	Data *bytes.Buffer
}

// Row is the dmarc row in a report
type Row struct {
	SourceIP        string
	Count           int64
	EvalDisposition string
	EvalSPFAlign    string
	EvalDKIMAalign  string
	Reason          string
	DKIMDomain      string
	DKIMResult      string
	SPFDomain       string
	SPFResult       string
	IdentifierHFrom string
}

// Rows is jus the report and the rows of a report
type Rows struct {
	Report Report
	Rows   []Row
}

// Report is the content of the report
type Report struct {
	ID                     int64
	ReportBegin            time.Time
	ReportEnd              time.Time
	PolicyDomain           string
	ReportOrg              string
	ReportID               string
	ReportEmail            string
	ReportExtraContactInfo string
	PolicyAdkim            string
	PolicyAspf             string
	PolicyP                string
	PolicySP               string
	PolicyPCT              string
	Count                  int64
	DKIMResult             string
	SPFResult              string
	Items                  int
}

// Reports is the collection of reports
type Reports struct {
	Reports    []Report
	LastPage   int
	CurPage    int
	NextPage   int
	TotalPages int
	Pages      []int
}

type dateRange struct {
	XMLName xml.Name `xml:"date_range"`
	Begin   int64    `xml:"begin"`
	End     int64    `xml:"end"`
}

type reportMetadata struct {
	XMLName          xml.Name  `xml:"report_metadata"`
	OrgName          string    `xml:"org_name"`
	Email            string    `xml:"email"`
	ExtraContactInfo string    `xml:"extra_contact_info,omitempty"`
	ReportID         string    `xml:"report_id"`
	DateRange        dateRange `xml:"date_range"`
}

type policyPublished struct {
	XMLName xml.Name `xml:"policy_published"`
	Domain  string   `xml:"domain"`
	ADKIM   string   `xml:"adkim"`
	ASPF    string   `xml:"aspf"`
	P       string   `xml:"p"`
	SP      string   `xml:"sp"`
	PCT     string   `xml:"pct"`
}

type reason struct {
	XMLName xml.Name `xml:"reason"`
	Type    string   `xml:"type"`
	Comment string   `xml:"comment"`
}

type policyEvaluated struct {
	XMLName     xml.Name `xml:"policy_evaluated"`
	Disposition string   `xml:"disposition"`
	DKIM        string   `xml:"dkim"`
	SPF         string   `xml:"spf"`
	Reasons     []reason `xml:"reason"`
}

type row struct {
	XMLName         xml.Name        `xml:"row"`
	SourceIP        string          `xml:"source_ip"`
	Count           int64           `xml:"count"`
	PolicyEvaluated policyEvaluated `xml:"policy_evaluated"`
}

type identify struct {
	XMLName    xml.Name `xml:"identifiers"`
	HeaderFrom string   `xml:"header_from"`
}

type spf struct {
	XMLName xml.Name `xml:"spf"`
	Result  string   `xml:"result"`
}

type dkim struct {
	XMLName xml.Name `xml:"dkim"`
	Result  string   `xml:"result"`
}

type authResult struct {
	XMLName xml.Name `xml:"auth_results"`
	SPF     []spf    `xml:"spf"`
	DKIM    []dkim   `xml:"dkim"`
}

type record struct {
	XMLName     xml.Name   `xml:"record"`
	Rows        []row      `xml:"row"`
	Identifiers identify   `xml:"identifiers"`
	AuthResults authResult `xml:"auth_results"`
}

// Feedback contains the reports and file information
type Feedback struct {
	XMLName         xml.Name `xml:"feedback"`
	FromFile        string
	ReportMetadata  reportMetadata  `xml:"report_metadata"`
	PolicyPublished policyPublished `xml:"policy_published"`
	Records         []record        `xml:"record"`
}
