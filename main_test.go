//go:build !nots

package main

import (
	"os"
	// "regexp"
	"testing"

	// "github.com/datacharmer/cliptool/cmd"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestInjectedVersionIsGitTag(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata",
		RequireExplicitExec: true,
		Setup: func(env *testscript.Env) error {
			env.Setenv("git_home", os.Getenv("PWD"))
			return nil
		},
	})
}

/*
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
*/
