package config

import (
	"github.com/kubeshop/testkube/pkg/oauth"
	"golang.org/x/oauth2"
)

type Data struct {
	TelemetryEnabled bool       `json:"telemetryEnabled,omitempty"`
	Namespace        string     `json:"namespace,omitempty"`
	Initialized      bool       `json:"initialized,omitempty"`
	APIURI           string     `json:"apiURI,omitempty"`
	OAuth2Data       OAuth2Data `json:"oauth2Data"`
}

func (c *Data) EnableAnalytics() {
	c.TelemetryEnabled = true
}

func (c *Data) DisableAnalytics() {
	c.TelemetryEnabled = false
}

func (c *Data) SetNamespace(ns string) {
	c.Namespace = ns
}

func (c *Data) SetInitialized() {
	c.Initialized = true
}

// OAuth2Data contains oauth credentials
type OAuth2Data struct {
	Enabled      bool               `json:"enabled,omitempty"`
	Token        *oauth2.Token      `json:"token,omitempty"`
	ClientID     string             `json:"clientID,omitempty"`
	ClientSecret string             `json:"clientSecret,omitempty"`
	Provider     oauth.ProviderType `json:"providerType,omitempty"`
	Scopes       []string           `json:"scopes"`
}

// EnableOAuth is oauth enable method
func (c *Data) EnableOAuth() {
	c.OAuth2Data.Enabled = true
}

// DisableOauth is oauth disable method
func (c *Data) DisableOauth() {
	c.OAuth2Data.Enabled = false
}
