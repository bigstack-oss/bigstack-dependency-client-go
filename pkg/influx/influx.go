package influx

import (
	"context"
	"sync"

	influxv2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var (
	helper *Helper
	once   sync.Once
)

type ConnectClient interface {
	QueryAPI(string) api.QueryAPI
	Close()
}

type QueryApiClient interface {
	Query(context.Context, string) (*api.QueryTableResult, error)
}

type Helper struct {
	ConnectClient
	QueryApiClient
	Options
}

func initOptions(opts []Option) *Options {
	options := &Options{Auth: Auth{}}
	for _, o := range opts {
		o(options)
	}

	return options
}

func NewHelper(opts ...Option) (*Helper, error) {
	initedOpts := initOptions(opts)

	h := &Helper{Options: *initedOpts}
	h.ConnectClient = influxv2.NewClient(h.Options.Url, h.Options.Auth.Token)
	h.QueryApiClient = h.ConnectClient.QueryAPI(h.Options.Org)

	return h, nil
}

func NewGlobalHelper(opts ...Option) error {
	var h *Helper
	var err error

	once.Do(func() {
		h, err = NewHelper(opts...)
		if err != nil {
			return
		}

		helper = h
	})
	if err != nil {
		return err
	}

	return nil
}

func GetGlobalHelper() *Helper {
	return helper
}

func (h *Helper) Close() {
	h.ConnectClient.Close()
}
