package sinch

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Send Handles Post requests to the sinch api... returns nil if successful
func Send(url string, token string, body io.Reader) (*http.Response, error) {

	//jsonData, err := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Post call failed: %w", err)
	}

	log.Infof("Response body: %v", resp.Body)

	return resp, nil
}

/*
POST \
     -H "Authorization: Bearer {token}" \
     -H "Content-Type: application/json"  -d '
      {
          "from": "12345",
          "to": [
              "123456789"
          ],
          "body": "Hi there! How are you?"
      }' \
  "https://eu.sms.api.sinch.com/xms/v1/{service_plan_id}/batches"
*/
