package util

import (
	"os"

	"github.com/go-rod/rod"
	"gopkg.in/yaml.v3"
)

const CONFIG_FILE = "config.yml"

type SubscriptionAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Server struct {
	Country string `yaml:"country"`
	City    string `yaml:"city"`
}

type Config struct {
	Username      string  `yaml:"username"`
	Password      string  `yaml:"password"`
	UUID          string  `yaml:"uuid"`
	Server        *Server `yaml:"server"`
	Subscription  *SubscriptionAuth
	Device        string `yaml:"device"`
	WireguardFile string `yaml:"wireguardFile"`
}

func (sub *SubscriptionAuth) GetEncryptPassword(page *rod.Page, token string) error {
	res, err := page.Eval(`
		async (username, authorization) => {
			const res = await fetch(
				"/v2/api/wireguard/get-encrypt-password",
				{
					method: "POST",
					body: new URLSearchParams({username}).toString(),
					headers: {
						'content-type': 'application/x-www-form-urlencoded',
						accept: 'application/json',
						authorization
					}
				}
			)
		if (!res.ok) {
			throw Error(await res.text())
			return
		}
		const json = await res.json()
		if (!json.status) {
			throw Error(json.message)
		}
		return json.body.encrypPass
	}`, sub.Username, "Bearer "+token)
	if err == nil {
		sub.Password = res.Value.Str()
	}
	return err
}

func (conf *Config) Save() error {
	return WriteConfig(conf)
}

func WriteConfig(conf *Config) error {
	data, err := yaml.Marshal(&conf)
	if err == nil {
		err = os.WriteFile(CONFIG_FILE, data, 0644)
	}
	return err
}

func ReadConfig() (*Config, error) {
	conf := Config{}

	data, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
