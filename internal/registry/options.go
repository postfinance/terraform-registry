package registry

// Option function to configure the Terraform registry
type Option func(*Registry)

// WithHTTPListen configures the http listen address.
func WithHTTPListen(listenAddr string) Option {
	return func(reg *Registry) {
		reg.listenAddr = listenAddr
	}
}
