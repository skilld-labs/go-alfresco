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
	return client, client.Login(cfg.Username, cfg.Password)
}

func NewDAPI(c *rest.Client) DAPI {
	return DAPI{Client: c}
}
