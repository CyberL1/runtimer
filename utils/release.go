package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLatestVersion() (*LatestRelease, error) {
	resp, err := http.Get("https://github.com/CyberL1/runtimer/releases/latest")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	latest := &LatestRelease{}
	err = json.Unmarshal(body, latest)
	if err != nil {
		return nil, err
	}
	return latest, nil
}