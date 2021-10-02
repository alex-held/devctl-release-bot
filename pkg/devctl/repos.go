package devctl

import "os"

const (
	devctlIndexRepoName  = "devctl-index"
	devctlIndexRepoOwner = "kubernetes-sigs"
)

//GetDevctlIndexRepoName returns the devctl-index repo name
func GetDevctlIndexRepoName() string {
	override := os.Getenv("UPSTREAM_KREW_INDEX_REPO_NAME")
	if override != "" {
		return override
	}

	return devctlIndexRepoName
}

//GetDevctlIndexRepoOwner returns the devctl-index repo owner
func GetDevctlIndexRepoOwner() string {
	override := os.Getenv("UPSTREAM_KREW_INDEX_REPO_OWNER")
	if override != "" {
		return override
	}

	return devctlIndexRepoOwner
}
