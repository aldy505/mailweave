package tlsrpt_test

import (
	"os"
	"path"
	"testing"

	"github.com/aldy505/mailweave/tlsrpt"
)

func TestParse(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("google.com json", func(t *testing.T) {
		file, err := os.Open(path.Join(pwd, "../testdata/tlsrpt/google.com!example.com!1747094400!1747180799!001.json"))
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		report, err := tlsrpt.ParseReport(file, tlsrpt.CompressionTypeNone)
		if err != nil {
			t.Fatal(err)
		}

		if report.OrganizationName != "Google Inc." {
			t.Errorf("OrganizationName: got %s, want %s", report.OrganizationName, "Google Inc.")
		}
	})

	t.Run("google.com gzip", func(t *testing.T) {
		file, err := os.Open(path.Join(pwd, "../testdata/tlsrpt/google.com!example.com!1747094400!1747180799!001.json.gz"))
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		report, err := tlsrpt.ParseReport(file, tlsrpt.CompressionTypeGZIP)
		if err != nil {
			t.Fatal(err)
		}

		if report.OrganizationName != "Google Inc." {
			t.Errorf("OrganizationName: got %s, want %s", report.OrganizationName, "Google Inc.")
		}
	})
}
