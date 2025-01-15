package openstack

import (
	"os"

	"github.com/gophercloud/gophercloud/v2"
)

var (
	DefaultEndpointOpts = gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	}
)

type Options struct {
	ConfFile         string `yaml:"confFile"`
	IdentityEndpoint string `yaml:"identityEndpoint"`
	Auth             `yaml:"auth"`

	Domain  `yaml:"domain"`
	Tenant  `yaml:"tenant"`
	Project `yaml:"project"`
	User    `yaml:"user"`

	Password string `yaml:"password"`
	Passcode string `yaml:"passcode"`

	IdentityAPIVersion string `yaml:"identityAPIVersion"`
	ImageAPIVersion    string `yaml:"imageAPIVersion"`

	Scope *gophercloud.AuthScope `yaml:"scope"`
}

type Auth struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type Tenant struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Domain `yaml:"domain"`
}

type Project struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Domain `yaml:"domain"`
}

type User struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Domain `yaml:"domain"`
}

type Domain struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

func ConfFile(confFile string) Option {
	return func(o *Options) {
		o.ConfFile = confFile
	}
}

func AuthType(authType string) Option {
	return func(o *Options) {
		o.Auth.Type = authType
	}
}

func AuthUrl(AuthUrl string) Option {
	return func(o *Options) {
		o.Auth.Url = AuthUrl
	}
}

func UserID(userID string) Option {
	return func(o *Options) {
		o.User.ID = userID
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.User.Name = username
	}
}

func Password(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}

func Passcode(passcode string) Option {
	return func(o *Options) {
		o.Passcode = passcode
	}
}

func TenantID(tenantID string) Option {
	return func(o *Options) {
		o.Tenant.ID = tenantID
	}
}

func TenantName(tenantName string) Option {
	return func(o *Options) {
		o.Tenant.Name = tenantName
	}
}

func ProjectName(projectName string) Option {
	return func(o *Options) {
		o.Project.Name = projectName
	}
}

func DomainID(domainID string) Option {
	return func(o *Options) {
		o.Domain.ID = domainID
	}
}

func DomainName(domainName string) Option {
	return func(o *Options) {
		o.Domain.Name = domainName
	}
}

func ProjectDomainName(projectDomainName string) Option {
	return func(o *Options) {
		o.Project.Domain.Name = projectDomainName
	}
}

func UserDomainName(userDomainName string) Option {
	return func(o *Options) {
		o.User.Domain.Name = userDomainName
	}
}

func IdentityAPIVersion(identityAPIVersion string) Option {
	return func(o *Options) {
		o.IdentityAPIVersion = identityAPIVersion
	}
}

func ImageAPIVersion(imageAPIVersion string) Option {
	return func(o *Options) {
		o.ImageAPIVersion = imageAPIVersion
	}
}

func Scope(scope *gophercloud.AuthScope) Option {
	return func(o *Options) {
		o.Scope = scope
	}
}
