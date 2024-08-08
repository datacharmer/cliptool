//go:build nots

package main

/*
To execute the tests in this file, run the following:

export GIT_COMMIT="$(git describe --tags)"
go test \
  -ldflags "-X 'github.com/datacharmer/cliptool/cmd.CliptoolVersion=$(git describe --tags)'" \
  --tags nots
*/

import (
	"os"
	"regexp"
	"testing"

	"github.com/datacharmer/cliptool/cmd"
)

func TestVersion(t *testing.T) {
	setVersion := os.Getenv("GIT_COMMIT")
	if setVersion == "" {
		t.Fatal("variable GIT_COMMIT not set")
	}
	if setVersion != cmd.CliptoolVersion {
		t.Fatalf("expected version: '%s' - got: '%s'\n", setVersion, cmd.CliptoolVersion)
	}

}

func TestVersionRegexp(t *testing.T) {
	versionRegexp := regexp.MustCompile(`v\d+\.\d+\.\d+`)
	if !versionRegexp.Match([]byte(cmd.CliptoolVersion)) {
		t.Fatalf(" cmd.CliptoolVersion '%s' doesn't match the format 'v#.#.#'\n", cmd.CliptoolVersion)
	}
}
