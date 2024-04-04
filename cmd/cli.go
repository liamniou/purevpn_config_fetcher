package purevpnwg

import "github.com/Rikpat/purevpnwg/pkg/util"

type Context struct {
	Debug  bool
	Config *util.Config
}

var CLI struct {
	Debug  bool         `help:"Enable debug mode."`
	Config *util.Config `embed:"" envprefix:"PUREVPN_"`

	Login  LoginCmd  `cmd:"" help:"Login and store cookies."`
	Update UpdateCmd `cmd:"" help:"Updates wireguard config file."`
	Full   FullCmd   `cmd:"" help:"Runs full process (login, wireguard update) without creating cookies file and writing config (for example for docker run)"`
}
