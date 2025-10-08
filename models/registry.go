package models

import (
	"encoding/json"
	"io"
	"net/http"
	"tag-to-sha/config"
	"time"
)

type Registry struct {
	Url        string
	AuthUrl    string
	Token      string
	ValidToken bool
	AuthReq    bool
}

type LRegistry struct {
	L []Registry
}

func NewLRegistry(images []*Image, cfg *config.Config) *LRegistry {
	resp := LRegistry{
		L: []Registry{},
	}
	registryMap := make(map[string]bool)
	for _, image := range images {
		registryMap[image.Registry] = true
	}
	for url := range registryMap {
		validToken := cfg.Registries[url].Token != ""
		reg := Registry{
			Url:        url,
			Token:      cfg.Registries[url].Token,
			AuthUrl:    cfg.Registries[url].AuthUrl,
			AuthReq:    cfg.Registries[url].AuthReq,
			ValidToken: validToken,
		}
		resp.L = append(resp.L, reg)
	}

	return &resp
}

func (r *LRegistry) GetToken() error {
	for index, reg := range r.L {
		if reg.AuthReq && !reg.ValidToken && reg.AuthUrl != "" {
			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			resp, err := client.Get(reg.AuthUrl)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			var responseData struct {
				Token string `json:"token"`
			}
			if err := json.Unmarshal(body, &responseData); err != nil {
				return err
			}
			r.L[index].Token = responseData.Token
			r.L[index].ValidToken = true
		}
	}
	return nil
}
