package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Strum355/log"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type ConsulService struct {
	Service
	client *api.Client
	ID     string
}

func (c *ConsulService) Setup() error {
	config := api.Config{
		Address: viper.GetString("cloud.consul.host"),
		Token:   viper.GetString("cloud.consul.token"),
	}
	client, err := api.NewClient(&config)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *ConsulService) GetSharedSecret() error {
	fn := func() error {
		path := "cloud-token"
		kv, _, err := c.client.KV().Get(path, &api.QueryOptions{})
		if err != nil {
			return err
		}

		if kv == nil {
			return errors.New(fmt.Sprintf("key %s not set", path))
		}

		viper.Set("cloud.http.token", string(kv.Value))
		return nil
	}

	count := 4
	var err error
	for ; count > 0; count-- {
		err = fn()
		if err == nil {
			return nil
		}
		log.WithFields(log.Fields{
			"limit": 4,
			"count": count,
		}).WithError(err).Error("failed to get shared secret")
		time.Sleep(time.Second * 3)
	}
	return err
}
