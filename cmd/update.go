package purevpn_wg

import (
	"dev.azure.com/Rikpat/Home/_git/purevpn_wg/pkg/purevpn"
	"dev.azure.com/Rikpat/Home/_git/purevpn_wg/pkg/util"
	"dev.azure.com/Rikpat/Home/_git/purevpn_wg/pkg/wireguard"
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

	config, err := util.ReadConfig()
	if err != nil {
		return err
	}
	token, err := purevpn.GetToken(page, config.UUID)

	if err != nil {
		return err
	}

	server, err := purevpn.GetWireguardServer(page, config, token)
	if err == nil {
		err = wireguard.UpdateConfig([]byte(server))
	}
	return err
}
