package purevpnwg

import (
	"github.com/Rikpat/purevpnwg/pkg/purevpn"
	"github.com/Rikpat/purevpnwg/pkg/wireguard"
)

type UpdateCmd struct {
}

func (r *UpdateCmd) Run(ctx *Context) error {
	page, err := purevpn.InitPage()
	if err != nil {
		if page != nil {
			page.MustClose()
		}
		return err
	}
	defer page.MustClose()

	token, err := purevpn.GetToken(page, ctx.Config.UUID)

	if err != nil {
		return err
	}

	server, err := purevpn.GetWireguardServer(page, ctx.Config, token)
	if err == nil {
		err = wireguard.UpdateConfig([]byte(server), ctx.Config)
	}
	return err
}
