package wireguard

import (
	"errors"
	"os"

	"dev.azure.com/Rikpat/Home/_git/purevpn_wg/pkg/util"
	"gopkg.in/ini.v1"
)

type WireguardConfig struct {
	Interface
	Peer
}

type Interface struct {
	PrivateKey, Address, DNS, Table string
}

type Peer struct {
	PublicKey, AllowedIPs, Endpoint, PersistentKeepalive string
}

func UpdateConfig(newConfig []byte) error {
	config, err := util.ReadConfig()
	if err != nil {
		return err
	}
	if _, err := os.Stat(config.WireguardFile); errors.Is(err, os.ErrNotExist) {
		wgConfFile, err := ini.Load(newConfig)
		if err != nil {
			return err
		}
		wgConfFile.SaveTo(config.WireguardFile)
		return nil
	}
	wgConf := new(WireguardConfig)

	err = ini.MapTo(wgConf, newConfig)
	if err != nil {
		return err
	}

	wgConfFile, err := ini.Load(config.WireguardFile)
	if err != nil {
		return err
	}
	iface := wgConfFile.Section("Interface")
	iface.Key("Address").SetValue(wgConf.Interface.Address)
	iface.Key("PrivateKey").SetValue(wgConf.Interface.PrivateKey)
	iface.Key("DNS").SetValue(wgConf.Interface.DNS)

	peer := wgConfFile.Section("Peer")
	peer.Key("PublicKey").SetValue(wgConf.Peer.PublicKey)
	peer.Key("Endpoint").SetValue(wgConf.Peer.Endpoint)

	return wgConfFile.SaveTo(config.WireguardFile)
}
