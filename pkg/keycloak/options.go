package keycloak

type Option func(*Options)

type Options struct {
	Host     string `json:"host" yaml:"host"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Realm    string `json:"realm" yaml:"realm"`
	Insecure bool   `json:"insecure" yaml:"insecure"`
}

func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
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

func Realm(realm string) Option {
	return func(o *Options) {
		o.Realm = realm
	}
}

func Insecure(insecure bool) Option {
	return func(o *Options) {
		o.Insecure = insecure
	}
}
