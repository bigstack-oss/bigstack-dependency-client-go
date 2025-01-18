package influx

var (
	Opts *Options
)

type Option func(*Options)

type Options struct {
	Url                   string `json:"url" yaml:"url"`
	Org                   string `json:"org" yaml:"org"`
	Auth                  `json:"auth" yaml:"auth"`
	TlsInsecureSkipVerify bool `json:"tlsInsecureSkipVerify" yaml:"tlsInsecureSkipVerify"`
	Timeout               uint `json:"timeout" yaml:"timeout"`
}

type Auth struct {
	Token string `json:"token" yaml:"token"`
}

func Url(url string) Option {
	return func(o *Options) {
		o.Url = url
	}
}

func Org(org string) Option {
	return func(o *Options) {
		o.Org = org
	}
}

func AuthToken(token string) Option {
	return func(o *Options) {
		o.Token = token
	}
}

func TlsInsecureSkipVerify(skip bool) Option {
	return func(o *Options) {
		o.TlsInsecureSkipVerify = skip
	}
}
