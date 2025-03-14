package main

import (
	"os"

	"github.com/owncloud/ocis-ocs/pkg/command"
)

func main() {
	if err := command.Root().Execute(); err != nil {
		os.Exit(1)
	}
}
