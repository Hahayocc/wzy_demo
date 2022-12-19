package base

import (
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/cache"
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/http"
	"errors"
	"log"
)

type Client interface {
	Get(key string) (string, error)
	GetKeys(keys ...string) (map[string]string, error)
	UpdateCache() error
}

type internalClient struct {
	cache    *cache.Cache
	ccClient *http.Client
	ticker   *cache.Ticker
}

func create() *internalClient {
	return &internalClient{}
}

func Start() (Client, error) {
	config := NewClientConfig()
	return StartWithConfig(config)
}

func StartWithConfig(config *clientConfig) (Client, error) {
	client := create()

	client.ccClient = http.NewClient(func(options *http.Options) {
		options.Timeout = config.GetTimeout()
	})
	client.cache = cache.NewCache()

	// first update cache
	err := cache.UpdateCache(client.cache, client.ccClient)
	if err != nil {
		log.Printf("first update cache err, err = %v", err)
		return nil, err
	}

	// start ticker
	client.ticker = cache.NewTicker(client.cache, client.ccClient, config.GetUpdateInterval())

	log.Println("start finished!")

	return client, nil
}

func (c *internalClient) Get(key string) (string, error) {
	v, _, err := c.getWithCache(key)
	if err != nil {
		return "", err
	}
	value := v.Object.(string)
	return value, nil
}

func (c *internalClient) GetKeys(keys ...string) (map[string]string, error) {
	kvs := make(map[string]string, len(keys))
	for _, k := range keys {
		v, _, err := c.getWithCache(k)
		if err != nil {
			return nil, err
		}
		value := v.Object.(string)
		kvs[k] = value
	}
	return kvs, nil
}

func (c *internalClient) UpdateCache() error {
	err := cache.UpdateCache(c.cache, c.ccClient)
	if err != nil {
		log.Printf("update cache err, err = %v", err)
		return err
	}
	return nil
}

func (c *internalClient) getWithCache(key string) (*cache.Item, bool, error) {
	item, exist := c.cache.Get(key)
	if !exist {
		return nil, false, errors.New("item not exist")
	}
	return item, true, nil

}
