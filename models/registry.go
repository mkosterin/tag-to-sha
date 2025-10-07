package models

type Registry struct {
	Url     string
	Token   string
	AuthReq bool
}

func NewRegistryList(images []*Image) []Registry {
	registryMap := make(map[string]bool)
	for _, image := range images {
		registryMap[image.Registry] = true
	}
	var registries []Registry
	for url := range registryMap {
		reg := Registry{
			Url:     url,
			Token:   "",
			AuthReq: true,
		}
		registries = append(registries, reg)
	}

	return registries
}
