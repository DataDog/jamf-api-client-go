package classic

type Option func(*Options) error
type Options struct {
	useTokenAuth bool
}

func resolveOptions(opts []Option) (*Options, error) {
	o := &Options{
		useTokenAuth: false,
	}
	for _, option := range opts {
		err := option(o)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func WithTokenAuth() Option {
	return func(o *Options) error {
		o.useTokenAuth = true
		return nil
	}
}
