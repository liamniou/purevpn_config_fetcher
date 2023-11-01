package util

import (
	"encoding/gob"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod/lib/proto"
)

func FilterCookies(cookies []*proto.NetworkCookie, namePrefix string) (filtered []*proto.NetworkCookie) {
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, namePrefix) {
			filtered = append(filtered, cookie)
		}
	}
	return filtered
}

func AreCookiesExpired(cookies []*proto.NetworkCookie) bool {
	now := proto.TimeSinceEpoch(time.Now().Unix())
	for _, cookie := range cookies {
		if cookie.Expires != -1 && cookie.Expires < now {
			return true
		}
	}
	return false
}

func WriteCookies(filePath string, cookies []*proto.NetworkCookie) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(cookies)
	}
	file.Close()
	return err
}

func ReadCookies(filePath string) ([]*proto.NetworkCookie, error) {
	cookies := []*proto.NetworkCookie{}
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&cookies)
	}
	file.Close()
	return cookies, err
}
