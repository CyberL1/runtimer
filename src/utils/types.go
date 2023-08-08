package utils

type GithubRelease struct {
	TagName    string `json:"tag_name"`
	Prerelease bool
}

type Runtime struct {
	Name string `json:"name"`
	Directory string `json:"directory"`
	Files []GithubFile
}

type GithubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	DownloadUrl string `json:"download_url"`
}
