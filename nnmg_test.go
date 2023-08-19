package main

import (
	"os"
	"testing"

	"github.com/netterminalmachine/nano-migrate/cmd/nnmg"
	"github.com/rogpeppe/go-internal/testscript"
)

// build an executable just for testing. Named 'nn:mg' to minimize any possible collission with aliases in user space.
func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"nn:mg": nnmg.Main,
	}))
}

func TestCommands(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
	})
}
