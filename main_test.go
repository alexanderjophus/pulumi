package main

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestMain(t *testing.T) {
	gcpLocation := "europe-west2-a"
	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd),
		Config: map[string]string{
			"gcp:zone":    gcpLocation,
			"gcp:project": "fluent-webbing-218616",
		},
	})
}
