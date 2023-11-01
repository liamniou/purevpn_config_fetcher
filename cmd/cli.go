package purevpn_wg

type Context struct {
	Debug bool
}

var CLI struct {
	Debug bool `help:"Enable debug mode."`

	Login  LoginCmd  `cmd:"" help:"Login and store cookies."`
	Update UpdateCmd `cmd:"" help:"Updates wireguard config file."`
}
