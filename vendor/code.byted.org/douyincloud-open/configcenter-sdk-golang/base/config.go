package base

import (
	"errors"
	"log"
	"time"
)

const (
	DefaultUpdateInterval = 1 * time.Minute
	DefaultTimeout        = 1 * time.Second
)

type clientConfig struct {
	updateInterval time.Duration
	timeout        time.Duration
}

func NewClientConfig() *clientConfig {
	c := &clientConfig{
		updateInterval: DefaultUpdateInterval,
		timeout:        DefaultTimeout,
	}
	return c
}

func (c *clientConfig) SetUpdateInterval(t time.Duration) error {
	if t < 10*time.Second {
		err := errors.New("parameter setting out of range, " + "parameter = " + t.String())
		log.Printf("set update interval err, err = %v", err)
		return err
	}
	c.updateInterval = t
	return nil
}

func (c *clientConfig) SetTimeout(t time.Duration) {
	//TODO 超时时间的范围限制
	c.timeout = t
}

func (c *clientConfig) GetUpdateInterval() time.Duration {
	return c.updateInterval
}

func (c *clientConfig) GetTimeout() time.Duration {
	return c.timeout
}
