package main

import (
	"os"

	"github.com/chen7gx/cxspn/app"
	"github.com/chen7gx/cxspn/cmd/cxspnd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
