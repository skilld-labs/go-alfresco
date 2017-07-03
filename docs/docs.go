package docs

import (
	"../rest"
	"strings"
)

type Config struct {
	Endpoint string
	Username string
	Password string
}

type DAPI struct {
	Client *rest.Client
}

func New(cfg Config) (*rest.Client, error) {
	var endpoint = strings.Split(cfg.Endpoint, ":")
	client := rest.NewClient(endpoint[0], endpoint[1], false)
	return client, client.Login(rest.Credentials{Username: cfg.Username, Password: cfg.Password}, false)
}

func NewDAPI(c *rest.Client) DAPI {
	return DAPI{Client: c}
}
