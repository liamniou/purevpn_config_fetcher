package main

import (
	purevpn_wg "dev.azure.com/Rikpat/Home/_git/purevpn_wg/cmd"
	"dev.azure.com/Rikpat/Home/_git/purevpn_wg/pkg/util"
	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
)

func main() {
	ctx := kong.Parse(&purevpn_wg.CLI, kong.Configuration(kongyaml.Loader, "/etc/purevpn_wg/config.yml", "~/.purevpn_wg.yml", util.CONFIG_FILE))
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&purevpn_wg.Context{Debug: purevpn_wg.CLI.Debug})
	ctx.FatalIfErrorf(err)
}
