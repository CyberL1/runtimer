package utils

type GithubRelease struct {
	TagName    string `json:"tag_name"`
	Prerelease bool
}

type GithubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	GitUrl      string `json:"git_url"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
}
