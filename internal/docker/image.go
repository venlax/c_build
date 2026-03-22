package docker

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

var baseURL = "http://127.0.0.1:5000"

type response struct {
	Input string  `json:"input"`
	Match *string `json:"match"`
}

func GetTheOptimalImageTag(version string) string {
	tag, err := getTag(version)
	if err != nil {
		slog.Error("get version wrong")
		panic("get version wrong")
	}
	return tag
}

func getTag(version string) (string, error) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	u := fmt.Sprintf("%s/match?version=%s", baseURL, url.QueryEscape(version))

	resp, err := client.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}

	if r.Match == nil {
		return "", nil
	}

	return *r.Match, nil
}