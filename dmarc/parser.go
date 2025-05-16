package dmarc

import (
	"encoding/xml"
	"fmt"
	"io"
)

func ParseFeedback(r io.Reader) (Feedback, error) {
	var feedback Feedback
	err := xml.NewDecoder(r).Decode(&feedback)
	if err != nil {
		return Feedback{}, fmt.Errorf("failed to parse feedback: %w", err)
	}

	return feedback, nil
}
