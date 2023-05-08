//go:build cli

package main

import (
	"vita/transport/cli"
)

func main() {
	cli.RootCmd.Execute()
}
