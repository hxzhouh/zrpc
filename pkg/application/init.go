package application

import (
	"github.com/hxzhouh/zrpc/pkg"
	"github.com/hxzhouh/zrpc/pkg/flag"
	"os"
)

func init() {
	flag.Register(&flag.BoolFlag{
		Name:    "version",
		Usage:   "--version, print version",
		Default: false,
		Action: func(string, *flag.FlagSet) {
			pkg.PrintVersion()
			os.Exit(0)
		},
	})
}
