// Package provider implements the backends for the Terraform provider registry
package provider

// NewProvider returns a new provider backend
// bc represents the backend client
func NewProvider(bc interface{}) (provider.Backend, error) {
	switch v := bc.(Type) {
	case artifactory.Client:
		return apb.New(v)
	}
	return nil, err
}
