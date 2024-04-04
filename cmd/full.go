package purevpnwg

import (
	"fmt"

	"github.com/Rikpat/purevpnwg/pkg/purevpn"
	"github.com/Rikpat/purevpnwg/pkg/util"
	"github.com/Rikpat/purevpnwg/pkg/wireguard"
)

type FullCmd struct {
	Config *util.Config `embed:"" envprefix:"PUREVPN_"`
}

func (r *FullCmd) Run(ctx *Context) error {
	page, cookies := purevpn.Login(r.Config.Username, r.Config.Password)
	defer page.MustClose()

	if ctx.Debug {
		fmt.Println("Successfully Logged in")
	}

	token := util.FilterCookies(cookies, "fa_token")
	if len(token) == 0 {
		return fmt.Errorf("no token in cookies")
	}

	userData, err := purevpn.GetUserData(token[0].Value)
	if err != nil {
		return err
	}

	if ctx.Debug {
		fmt.Println("Successfully parsed user data")
	}

	if r.Config.Subscription.Username == "" {
		if r.Config.Subscription, err = userData.SelectSubscription(); err != nil {
			return err
		}
	}

	if ctx.Debug {
		fmt.Printf("Selected subscription %v\n", r.Config.Subscription.Username)
	}

	if err := r.Config.Subscription.GetEncryptPassword(page, token[0].Value); err != nil {
		return err
	}

	if ctx.Debug {
		fmt.Println("Successfully got subscription password")
	}

	server, err := purevpn.GetWireguardServer(page, r.Config, token[0].Value)

	if err == nil {
		if ctx.Debug {
			fmt.Printf("Got wireguard server: %v\n", server)
		}
		err = wireguard.UpdateConfig([]byte(server), r.Config)
		if ctx.Debug && err == nil {
			fmt.Printf("Created wireguard file at %v\n", r.Config.WireguardFile)
		}
	}
	return err
}
