package models

type Registry struct {
	url   string
	token string
}

func NewRegistry() (*Registry, error) {
	return &Registry{
		url:   "",
		token: "",
	}, nil
}

func (r *Registry) GetToken() error {

	return nil
}
