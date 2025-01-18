package openstack

import (
	"os"

	"github.com/gophercloud/gophercloud"
)

var (
	DefaultEndpointOpts = gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	}
)

type Options struct {
	ConfFile         string `json:"confFile" yaml:"confFile"`
	IdentityEndpoint string `json:"identityEndpoint" yaml:"identityEndpoint"`
	AuthType         string `json:"authType" yaml:"authType"`

	UserID   string `json:"userID" yaml:"userID"`
	Username string `json:"username" yaml:"username"`

	Password string `json:"password" yaml:"password"`
	Passcode string `json:"passcode" yaml:"passcode"`

	TenantID    string `json:"tenantID" yaml:"tenantID"`
	TenantName  string `json:"tenantName" yaml:"tenantName"`
	ProjectName string `json:"projectName" yaml:"projectName"`

	DomainID          string `json:"domainID" yaml:"domainID"`
	DomainName        string `json:"domainName" yaml:"domainName"`
	ProjectDomainName string `json:"projectDomainName" yaml:"projectDomainName"`
	UserDomainName    string `json:"userDomainName" yaml:"userDomainName"`

	IdentityAPIVersion string `json:"identityAPIVersion" yaml:"identityAPIVersion"`
	ImageAPIVersion    string `json:"imageAPIVersion" yaml:"imageAPIVersion"`

	Scope *gophercloud.AuthScope `json:"scope" yaml:"scope"`
}

func ConfFile(confFile string) Option {
	return func(o *Options) {
		o.ConfFile = confFile
	}
}

func IdentityEndpoint(identityEndpoint string) Option {
	return func(o *Options) {
		o.IdentityEndpoint = identityEndpoint
	}
}

func AuthType(authType string) Option {
	return func(o *Options) {
		o.AuthType = authType
	}
}

func UserID(userID string) Option {
	return func(o *Options) {
		o.UserID = userID
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.Username = username
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
		o.TenantID = tenantID
	}
}

func TenantName(tenantName string) Option {
	return func(o *Options) {
		o.TenantName = tenantName
	}
}

func ProjectName(projectName string) Option {
	return func(o *Options) {
		o.ProjectName = projectName
	}
}

func DomainID(domainID string) Option {
	return func(o *Options) {
		o.DomainID = domainID
	}
}

func DomainName(domainName string) Option {
	return func(o *Options) {
		o.DomainName = domainName
	}
}

func ProjectDomainName(projectDomainName string) Option {
	return func(o *Options) {
		o.ProjectDomainName = projectDomainName
	}
}

func UserDomainName(userDomainName string) Option {
	return func(o *Options) {
		o.UserDomainName = userDomainName
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
