package tlsrpt

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
)

type CompressionType uint8

const (
	CompressionTypeNone CompressionType = iota
	CompressionTypeGZIP
)

func ParseReport(r io.Reader, compression CompressionType) (*Report, error) {
	// decompress file if compression is gzip
	if compression == CompressionTypeGZIP {
		reader, err := gzip.NewReader(r)
		if err != nil {
			return nil, fmt.Errorf("decompressing gzip: %w", err)
		}
		defer reader.Close()

		r = reader
	}

	// parse report
	var report Report
	err := json.NewDecoder(r).Decode(&report)
	if err != nil {
		return nil, fmt.Errorf("parsing json report: %w", err)
	}

	return &report, nil
}
