package devctl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevctlIndexRepoName(t *testing.T) {
	testcases := []struct {
		name     string
		setup    func()
		expected string
	}{
		{
			name:     "env variable is not set",
			expected: "devctl-index",
		},
		{
			name: "env variable is set to empty value",
			setup: func() {
				os.Setenv("UPSTREAM_KREW_INDEX_REPO_NAME", "")
			},
			expected: "devctl-index",
		},
		{
			name: "env variable is set to value",
			setup: func() {
				os.Setenv("UPSTREAM_KREW_INDEX_REPO_NAME", "foo-bar")
			},
			expected: "foo-bar",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Clearenv()
			if tc.setup != nil {
				tc.setup()
			}

			actual := GetDevctlIndexRepoName()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetDevctlIndexRepoOwner(t *testing.T) {
	testcases := []struct {
		name     string
		setup    func()
		expected string
	}{
		{
			name:     "env variable is not set",
			expected: "kubernetes-sigs",
		},
		{
			name: "env variable is set to empty value",
			setup: func() {
				os.Setenv("UPSTREAM_KREW_INDEX_REPO_OWNER", "")
			},
			expected: "kubernetes-sigs",
		},
		{
			name: "env variable is set to value",
			setup: func() {
				os.Setenv("UPSTREAM_KREW_INDEX_REPO_OWNER", "foo-bar-user")
			},
			expected: "foo-bar-user",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Clearenv()
			if tc.setup != nil {
				tc.setup()
			}

			actual := GetDevctlIndexRepoOwner()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
