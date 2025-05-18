package dmarc_test

import (
	"os"
	"path"
	"testing"

	"github.com/aldy505/mailweave/dmarc"
)

func TestParseAmazonFeedback(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("amazon xml", func(t *testing.T) {
		f, err := os.Open(path.Join(pwd, "../testdata/dmarc/amazonses.com!example.com!1747180800!1747267200.xml"))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		feedback, err := dmarc.ParseFeedback(f)
		if err != nil {
			t.Fatal(err)
		}

		if feedback.ReportMetadata.OrgName != "AMAZON-SES" {
			t.Errorf("OrgName = %s, want AMAZON-SES", feedback.ReportMetadata.OrgName)
		}
		if feedback.ReportMetadata.Email != "postmaster@amazonses.com" {
			t.Errorf("Email = %s, want postmaster@amazonses.com", feedback.ReportMetadata.Email)
		}
	})
}
