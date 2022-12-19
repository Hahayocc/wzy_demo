package cache

import (
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/http"
	"code.byted.org/douyincloud-open/configcenter-sdk-golang/openapi"
	"context"
	"encoding/json"
	"log"
	"time"
)

type Ticker struct {
	StopChan chan bool
}

func NewTicker(cache *Cache, ccClient *http.Client, interval time.Duration) *Ticker {
	ticker := time.NewTicker(interval)
	stopChan := make(chan bool)

	go func(ticker *time.Ticker) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("error: %v", err)
			}
			ticker.Stop()
		}()
		for {
			select {
			case <-ticker.C:
				log.Printf("自动轮巡,当前时间: %s", time.Now())
				UpdateCache(cache, ccClient)
			case stop := <-stopChan:
				if stop {
					log.Println("Ticker Stop")
					return
				}
			}
		}
	}(ticker)

	return &Ticker{StopChan: stopChan}
}

func UpdateCache(cache *Cache, ccClient *http.Client) error {
	configVersion := cache.GetVersion()
	if configVersion == "" {
		configVersion = "0"
	}
	bodyStruct := openapi.GetConfigListReqBody{Version: configVersion}
	jsonByte, _ := json.Marshal(bodyStruct)
	body := string(jsonByte)

	//TODO: 优化入、出参
	respBody, _, _, _, err := ccClient.CtxHttpPostRaw(context.Background(), body, nil)
	if err != nil {
		return err
	}

	var resp openapi.GetConfigListResponse
	var httpResult openapi.HttpResp
	json.Unmarshal(respBody, &httpResult)
	resp = httpResult.Data

	if configVersion >= resp.Version {
		return nil
	}

	// 更新缓存
	items := make(map[string]*Item, len(resp.Kvs))
	for _, v := range resp.Kvs {
		items[v.Key] = &Item{
			Object: v.Value,
			Type:   v.Type,
		}
	}
	cache.Set(items)
	cache.SetVersion(resp.Version)
	return nil
}
