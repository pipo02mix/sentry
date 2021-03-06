package nslabels

import (
	"github.com/jasonrichardsmith/sentry/config"
	"github.com/jasonrichardsmith/sentry/sentry"
)

const (
	NAME = "nslabels"
)

func init() {
	config.Register(&Config{})
}

type Config struct {
	IgnoredNamespaces []string `mapstructure:"ignoredNamespaces"`
}

func (c *Config) Name() string {
	return NAME
}

func (c *Config) LoadSentry() sentry.Sentry {
	return Sentry{c.IgnoredNamespaces}
}
