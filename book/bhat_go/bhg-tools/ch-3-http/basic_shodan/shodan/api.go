package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	UnlockedLeft int    `json:"unclocked_left"`
	MonitoredIPS int    `json:"monitored_ips"`
}

func (s *Client) apiInfo() (*apiInfo, error) {
	res, err = http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKet))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret apiInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
