//go:build !nots

package main

import (
	"os"
	"testing"

	"github.com/datacharmer/cliptool/cmd"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestInjectedVersionIsGitTag(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata",
		RequireExplicitExec: true,
		Setup: func(env *testscript.Env) error {
			env.Setenv("git_home", os.Getenv("PWD"))
			env.Setenv("the_regexp_version", ".*"+cmd.CliptoolVersion)
			env.Setenv("the_exact_version", cmd.CliptoolVersion)
			return nil
		},
	})
}
