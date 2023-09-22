package main

import (
	"os"

	cdkvalidiumnode "github.com/0xPolygon/cdk-validium-node"
	"github.com/urfave/cli/v2"
)

func versionCmd(*cli.Context) error {
	cdkvalidiumnode.PrintVersion(os.Stdout)
	return nil
}
