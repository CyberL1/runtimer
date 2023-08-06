package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"runtimer/constants"
)

func GetLatestCliVersion() (*GithubRelease, error) {
	resp, err := http.Get(constants.GithubReleaseUrl)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	release := &GithubRelease{}
	err = json.Unmarshal(body, release)
	if err != nil {
		return nil, err
	}
	return release, nil
}
