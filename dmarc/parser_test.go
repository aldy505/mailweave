package dmarc_test

import (
	"os"
	"path"
	"testing"

	"github.com/aldy505/mailweave/dmarc"
)

func TestParseFeedback(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("google xml", func(t *testing.T) {
		f, err := os.Open(path.Join(pwd, "../testdata/dmarc/google.com!example.com!1747008000!1747094399"))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		feedback, err := dmarc.ParseFeedback(f)
		if err != nil {
			t.Fatal(err)
		}

		if feedback.ReportMetadata.OrgName != "google.com" {
			t.Errorf("OrgName = %s, want google.com", feedback.ReportMetadata.OrgName)
		}
		if feedback.ReportMetadata.Email != "noreply-dmarc-support@google.com" {
			t.Errorf("Email = %s, want noreply-dmarc-support@google.com", feedback.ReportMetadata.Email)
		}
		if feedback.ReportMetadata.ReportID != "8639335954371369510" {
			t.Errorf("ReportID = %s, want 8639335954371369510", feedback.ReportMetadata.ReportID)
		}
	})
}
