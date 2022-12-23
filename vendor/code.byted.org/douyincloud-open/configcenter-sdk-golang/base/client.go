package base

import (
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/cache"
	error2 "code.byted.org/douyincloud-open/configcenter-sdk-golang/error"
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/http"
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/openapi"
	"context"
	"encoding/json"
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

func Start() Client {
	config := NewClientConfig()
	return StartWithConfig(config)
}

func StartWithConfig(config *clientConfig) Client {
	client := create()

	client.ccClient = http.NewClient(func(options *http.Options) {
		options.Timeout = config.GetTimeout()
	})
	client.cache = cache.NewCache()

	// first update cache
	err := updateCache(client.cache, client.ccClient)
	if err != nil {
		log.Printf("first update cache err, err = %v", err)
	}

	// start ticker
	client.ticker = cache.NewTicker(client.cache, client.ccClient, config.GetUpdateInterval())

	log.Println("start finished!")

	return client
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
	err := updateCache(c.cache, c.ccClient)
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

func updateCache(cache2 *cache.Cache, ccClient *http.Client) error {
	configVersion := cache2.GetVersion()
	if configVersion == "" {
		configVersion = "0"
	}
	bodyStruct := openapi.GetConfigListReqBody{Version: configVersion}
	jsonByte, _ := json.Marshal(bodyStruct)
	body := string(jsonByte)

	//TODO: 优化入、出参
	respBody, _, err := ccClient.CtxHttpPostRaw(context.Background(), body, nil)
	if err != nil {
		return err
	}

	var resp openapi.GetConfigListResponse
	var httpResult openapi.HttpResp
	json.Unmarshal(respBody, &httpResult)
	resp = httpResult.Data
	code := httpResult.Code
	msg := httpResult.Msg

	if code != 0 {
		err := error2.NewErr(2, "request err", code, msg)
		return err
	}

	if configVersion >= resp.Version {
		return nil
	}

	// 更新缓存
	items := make(map[string]*cache.Item, len(resp.Kvs))
	for _, v := range resp.Kvs {
		items[v.Key] = &cache.Item{
			Object: v.Value,
			Type:   v.Type,
		}
	}
	cache2.Set(items)
	cache2.SetVersion(resp.Version)
	return nil
}
