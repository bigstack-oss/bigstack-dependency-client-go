package keycloak

type Option func(*Options)

type Options struct {
	Host                  string `json:"host" yaml:"host"`
	TlsInsecureSkipVerify bool   `json:"tlsInsecureSkipVerify" yaml:"tlsInsecureSkipVerify"`
	Auth                  `json:"auth" yaml:"auth"`
}

type Auth struct {
	Realm    string `json:"realm" yaml:"realm"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func Insecure(insecure bool) Option {
	return func(o *Options) {
		o.TlsInsecureSkipVerify = insecure
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.Auth.Username = username
	}
}

func Password(password string) Option {
	return func(o *Options) {
		o.Auth.Password = password
	}
}

func Realm(realm string) Option {
	return func(o *Options) {
		o.Auth.Realm = realm
	}
}
