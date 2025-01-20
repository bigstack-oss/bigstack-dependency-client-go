package openstack

import (
	"context"
	"os"
	"sync"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	log "go-micro.dev/v5/logger"
)

var (
	Opts   *Options
	helper *Helper

	once sync.Once
)

type Helper struct {
	Provider *gophercloud.ProviderClient

	Identity *gophercloud.ServiceClient
	Compute  *gophercloud.ServiceClient
	Network  *gophercloud.ServiceClient
	Storage  *gophercloud.ServiceClient
	Share    *gophercloud.ServiceClient

	*Options
}

type Option func(*Options)

func GetGlobalHelper() *Helper {
	return helper
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

func NewHelper(opts ...Option) (*Helper, error) {
	provider, err := newProvider(opts...)
	if err != nil {
		log.Errorf("failed to create provider: %s", err.Error())
		return nil, err
	}

	identityCli, err := newIdentityCli(provider)
	if err != nil {
		log.Errorf("failed to create identity client: %s", err.Error())
		return nil, err
	}

	computeCli, err := newComputeCli(provider)
	if err != nil {
		log.Errorf("failed to create compute client: %s", err.Error())
		return nil, err
	}

	networkCli, err := newNetworkCli(provider)
	if err != nil {
		log.Errorf("failed to create network client: %s", err.Error())
		return nil, err
	}

	storageCli, err := newStorageCli(provider)
	if err != nil {
		log.Errorf("failed to create storage client: %s", err.Error())
		return nil, err
	}

	shareCli, err := newShareCli(provider)
	if err != nil {
		log.Errorf("failed to create share client: %s", err.Error())
		return nil, err
	}

	return &Helper{
		Provider: provider,
		Identity: identityCli,
		Compute:  computeCli,
		Network:  networkCli,
		Storage:  storageCli,
		Share:    shareCli,
	}, nil
}

func newProvider(opts ...Option) (*gophercloud.ProviderClient, error) {
	syncedOpts, err := syncOptions(opts)
	if err != nil {
		return nil, err
	}

	finalOpts, err := genAuthOpts(syncedOpts)
	if err != nil {
		return nil, err
	}

	return openstack.AuthenticatedClient(
		context.Background(),
		finalOpts,
	)
}

func syncOptions(opts []Option) (*Options, error) {
	options, err := NewConf()
	if err != nil {
		return nil, err
	}

	for _, o := range opts {
		o(options)
	}

	return options, nil
}

func genAuthOpts(opts *Options) (gophercloud.AuthOptions, error) {
	if opts.Auth.Type == "env" {
		return openstack.AuthOptionsFromEnv()
	}

	return gophercloud.AuthOptions{
		IdentityEndpoint: opts.Auth.Url,
		Username:         opts.User.Name,
		Password:         opts.Password,
		TenantName:       opts.Project.Name,
		DomainName:       opts.Domain.Name,
	}, nil
}

func NewConf() (*Options, error) {
	opts := &Options{
		Domain: Domain{
			Name: "default",
		},
		Auth: Auth{
			Type: os.Getenv("OS_AUTH_TYPE"),
			Url:  os.Getenv("OS_AUTH_URL"),
		},
		User: User{
			Name: os.Getenv("OS_USERNAME"),
			Domain: Domain{
				Name: os.Getenv("OS_USER_DOMAIN_NAME"),
			},
		},
		Password: os.Getenv("OS_PASSWORD"),
		Tenant: Tenant{
			Name: os.Getenv("OS_PROJECT_NAME"),
			Domain: Domain{
				Name: os.Getenv("OS_PROJECT_DOMAIN_NAME"),
			},
		},
	}

	systemScope := os.Getenv("OS_SYSTEM_SCOPE")
	if systemScope == "all" {
		opts.Scope = &gophercloud.AuthScope{
			System: true,
		}
	}

	return opts, nil
}

func newIdentityCli(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(
		provider,
		gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		},
	)
}

func newComputeCli(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	return openstack.NewComputeV2(
		provider,
		gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		},
	)
}

func newNetworkCli(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	return openstack.NewNetworkV2(
		provider,
		gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		},
	)
}

func newStorageCli(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	return openstack.NewBlockStorageV3(
		provider,
		gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		},
	)
}

func newShareCli(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	return openstack.NewSharedFileSystemV2(
		provider,
		gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		},
	)
}
