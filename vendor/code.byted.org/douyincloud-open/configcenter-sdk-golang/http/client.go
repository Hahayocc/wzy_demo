package http

import (
	"context"
	"net/http"
)

const (
	HttpMethodPost = "POST"
)

type Client struct {
	opts *Options
	c    http.Client
}

func NewClient(options ...Option) *Client {
	opts := DefaultOptions()
	for _, o := range options {
		o(opts)
	}
	return &Client{
		opts: opts,
		c: http.Client{
			Transport: opts.Transport,
		},
	}
}

func (client *Client) GetOptions() *Options {
	return client.opts
}

type APIInfo struct {
	Method string
	Body   string
	Header http.Header
}

// HttpPostRaw 发起JSON的post请求
func (client *Client) HttpPostRaw(body string, headers http.Header) ([]byte, int, string, http.Header, error) {
	return client.CtxHttpPostRaw(context.Background(), body, headers)
}

func (client *Client) CtxHttpPostRaw(ctx context.Context, body string, headers http.Header) ([]byte, int, string, http.Header, error) {
	apiInfo := &APIInfo{
		Method: HttpMethodPost,
		Body:   body,
		Header: headers,
	}
	return client.request(ctx, apiInfo)
}
