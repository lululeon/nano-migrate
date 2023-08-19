package main

import (
	"os"

	"github.com/netterminalmachine/nano-migrate/cmd/nnmg"
)

func main() {
	status := nnmg.Main()
	os.Exit(status)
}
