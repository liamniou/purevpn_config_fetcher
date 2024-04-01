package purevpn_wg

import (
	"fmt"

	"github.com/Rikpat/purevpn_wg/pkg/purevpn"
	"github.com/Rikpat/purevpn_wg/pkg/util"
	"github.com/Rikpat/purevpn_wg/pkg/wireguard"
)

type FullCmd struct {
	Config *util.Config `embed:"" envprefix:"PUREVPN_"`
}

func (r *FullCmd) Run(ctx *Context) error {
	page, cookies := purevpn.Login(r.Config.Username, r.Config.Password)
	defer page.MustClose()

	token := util.FilterCookies(cookies, "fa_token")
	if len(token) == 0 {
		return fmt.Errorf("no token in cookies")
	}

	userData, err := purevpn.GetUserData(token[0].Value)
	if err != nil {
		return err
	}

	if r.Config.Subscription.Username == "" {
		if r.Config.Subscription, err = userData.SelectSubscription(); err != nil {
			return err
		}
	}

	if err := r.Config.Subscription.GetEncryptPassword(page, token[0].Value); err != nil {
		return err
	}

	server, err := purevpn.GetWireguardServer(page, r.Config, token[0].Value)
	if err == nil {
		err = wireguard.UpdateConfig([]byte(server), r.Config)
	}
	return err
}
