# Mailweave Development Guidelines

This document provides guidelines and information for developers working on the Mailweave project.

## Build/Configuration Instructions

### Local Development

1. **Prerequisites**:
   - Go 1.24 or later
   - Node.js 22.15 or later (for frontend development)

2. **Backend Build**:
   ```bash
   # Build the backend
   go build ./cmd/main.go
   
   # Run the backend
   ./main
   ```

3. **Frontend Build**:
   ```bash
   # Navigate to the frontend directory
   cd static
   
   # Install dependencies
   npm ci
   
   # Build the frontend
   npm run build
   ```

4. **Full Application Build**:
   ```bash
   # Build frontend
   cd static && npm ci && npm run build && cd ..
   
   # Build backend with version information
   go build -o mailweave -ldflags="-s -w -X 'main.version=$(git describe --tags --always)'" ./cmd/
   ```

### Docker Build

The project includes a multi-stage Dockerfile for containerized deployment:

```bash
# Build the Docker image
docker build -t mailweave .

# Run the container
docker run -p 8080:8080 mailweave
```

## Testing Information

### Running Tests

To run all tests in the project:

```bash
go test -v ./...
```

To run tests with coverage reporting:

```bash
go test -v -coverprofile=coverage.out -covermode=atomic ./...

# View coverage in browser
go tool cover -html=coverage.out
```

To run tests for a specific package:

```bash
go test -v ./dmarc
go test -v ./tlsrpt
go test -v ./datastore
```

### Adding New Tests

1. Create a test file with the naming convention `<filename>_test.go` in the same package as the code being tested.
2. For testing packages, use the `_test` suffix in the package name (e.g., `package dmarc_test`).
3. Use the standard Go testing package and follow the existing test patterns.

### Test Example

Here's an example of how to create a test for the DMARC parser:

```go
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
```

### Test Data

Sample DMARC and TLS-RPT reports for testing are stored in the `testdata` directory:
- `testdata/dmarc/` - Contains sample DMARC reports from various providers
- `testdata/tlsrpt/` - Contains sample TLS-RPT reports

When adding new tests, you can use these existing samples or add new ones as needed.

## Additional Development Information

### Code Structure

The project follows standard Go package organization:

- `cmd/` - Main application entry points
- `datastore/` - Database and storage interfaces
- `dmarc/` - DMARC report parsing and processing
- `tlsrpt/` - TLS-RPT report parsing and processing
- `static/` - Frontend code (React/TypeScript)
- `testdata/` - Test data files

### Code Style

- Follow standard Go code style and conventions
- Use `gofmt` or `goimports` to format code
- Write meaningful comments and documentation
- Include tests for new functionality

### Continuous Integration

The project uses GitHub Actions for CI/CD:
- Backend tests and builds
- Frontend builds
- Multi-architecture binary builds

The CI workflow is defined in `.github/workflows/ci.yml`.

### Dependencies

Dependencies are managed using Go modules. Add new dependencies with:

```bash
go get github.com/example/package
```

Frontend dependencies are managed with npm in the `static` directory.