package main

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestMain(t *testing.T) {
	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd),
		Config: map[string]string{
			"gcp:zone":    os.Getenv("GCP_ZONE"),
			"gcp:project": os.Getenv("GCP_PROJECT"),
		},
	})
}
