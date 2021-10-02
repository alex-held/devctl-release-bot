package devctl

import "os"

const (
	devctlIndexRepoName  = "devctl-index"
	devctlIndexRepoOwner = "kubernetes-sigs"
)

//GetKrewIndexRepoName returns the devctl-index repo name
func GetKrewIndexRepoName() string {
	override := os.Getenv("UPSTREAM_KREW_INDEX_REPO_NAME")
	if override != "" {
		return override
	}

	return devctlIndexRepoName
}

//GetKrewIndexRepoOwner returns the devctl-index repo owner
func GetKrewIndexRepoOwner() string {
	override := os.Getenv("UPSTREAM_KREW_INDEX_REPO_OWNER")
	if override != "" {
		return override
	}

	return devctlIndexRepoOwner
}
