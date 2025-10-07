package models

import (
	"log/slog"
	"strings"
)

type Image struct {
	Registry     string `json:"registry"`
	Path         string `json:"path"`
	Tag          string `json:"tag"`
	Sha256digest string `json:"sha256digest"`
}

func NewImage() (*Image, error) {
	return &Image{
		Registry:     "",
		Path:         "",
		Tag:          "",
		Sha256digest: "",
	}, nil
}

func (i *Image) ParseImage(line string, logger *slog.Logger) error {
	if strings.Contains(line, "@sha256:") {
		parts := strings.Split(line, "@sha256:")
		if len(parts) != 2 {
			logger.Error("error in image format with sha256", "line", line)
			return nil
		}
		i.Sha256digest = parts[1]
		line = parts[0]
	}

	registryAndPath := strings.Split(line, "/")
	if len(registryAndPath) == 1 || (!strings.Contains(registryAndPath[0], ".") && len(registryAndPath) == 2) {
		i.Registry = "docker.io"
		i.Path = strings.Join(registryAndPath, "/")

		if len(registryAndPath) == 1 {
			i.Path = "library/" + registryAndPath[0]
		}
	} else {
		i.Registry = registryAndPath[0]
		i.Path = strings.Join(registryAndPath[1:], "/")
	}

	if strings.Contains(i.Path, ":") {
		parts := strings.Split(i.Path, ":")
		i.Path = parts[0]
		i.Tag = parts[1]
	} else if i.Sha256digest == "" {
		i.Tag = "latest"
	}

	return nil
}
