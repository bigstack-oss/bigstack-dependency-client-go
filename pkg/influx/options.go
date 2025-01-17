package influx

var (
	Opts *Options
)

type Option func(*Options)

type Options struct {
	Url string `validate:"required"`
	Org string
	Auth
	TlsInsecureSkipVerify bool
	Timeout               uint
}

type Auth struct {
	Token string
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
