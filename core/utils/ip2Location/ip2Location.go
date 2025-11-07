package ip2Location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func IP2Location(ip string) *Location {
	return &Location{ip: ip}
}

func (root *Location) GetLocation() (string, error) {
	url := fmt.Sprintf("https://opendata.baidu.com/api.php?query=%v&resource_id=6006&oe=utf8", root.ip)
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("code not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data RootEntity

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if len(data.Data) == 0 {
		return "", nil
	}

	return data.Data[0].Location, nil
}
