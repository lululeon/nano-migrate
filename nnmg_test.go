package main

import (
	"os"
	"testing"

	"github.com/netterminalmachine/nano-migrate/cmd/nnmg"
	"github.com/rogpeppe/go-internal/testscript"
)

// build an executable just for testing
func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"nnmg": nnmg.Main,
	}))
}

func TestCommands(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
	})
}
