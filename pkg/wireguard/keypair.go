package wireguard

import "golang.zx2c4.com/wireguard/wgctrl/wgtypes"

func GenerateKeyPair() (string, string, error) {
	pk, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}
	return pk.String(), pk.PublicKey().String(), nil
}
