package sinch

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Send Handles Post requests to the sinch api... returns nil if successful
func Send(url string, body io.Reader) error {
	resp, err := http.Post(url, "Content-Type: application/json", body)

	if err != nil {
		return fmt.Errorf("Post call failed: %w", err)
	}

	log.Infof("Response body: %v", resp.Body)

	return nil
}
