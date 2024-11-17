package redis

import "github.com/zcolleen/shipper/internal"

type Client = internal.Client

type client interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

func NewClient() Client {
	return &clientEntity{}
}

type clientEntity struct {
}

func (c *clientEntity) Get(key string) (string, error) {
	return "", nil
}

func (c *clientEntity) Set(key, value string) error {
	return nil
}
