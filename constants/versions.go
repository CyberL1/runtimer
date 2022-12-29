package constants

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Version string

func GetLatestCliVersion() (*GithubRelease, error) {
	resp, err := http.Get(fmt.Sprintf("%s/releases/latest", GithubRepoApi))
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