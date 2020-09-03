package xlog_config

// WithEncoding ...
func WithEncoding(enc string) Option {
	return func(opts *config) {
		opts.Encoding = enc
	}
}

func WithLevel(ll string) Option {
	return func(opts *config) {
		opts.Level = ll
	}
}
