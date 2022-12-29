package constants

import "fmt"

var GithubUsername = "CyberL1"
var GithubRepo = "runtimer"

var GithubRepoApi = fmt.Sprintf("https://api.github.com/repos/%s/%s", GithubUsername, GithubRepo)