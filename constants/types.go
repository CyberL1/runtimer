package constants

type RuntimesType struct {
	Name string
	Url string
	Runtime string
	Version string
	Ext string
	Bin string
	Os map[string]OsType
	Arch map[string]string
}

type OsType struct {
	Name string
	Ext string
	Bin string
}

type GithubRelease struct {
	TagName string `json:"tag_name"`
	Prerelease bool
}