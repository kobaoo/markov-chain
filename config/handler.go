package config

import (
	"fmt"
	"strings"
)

func HandleConfig(cfg *Config) error {
	if cfg.MaxWords <= 0 || cfg.MaxWords > 10000 {
		return fmt.Errorf("length of maximum words must be between 1 and 10000")
	} else if cfg.PrefixLength < 1 || cfg.PrefixLength > 5 {
		return fmt.Errorf("prefix length must be between 1 and 5")
	}

	parts := strings.Split(cfg.StartPrefix, " ")
	if strings.TrimSpace(cfg.StartPrefix) != "" && len(parts) > cfg.PrefixLength {
		return fmt.Errorf("length of starting prefix (%s) is more than prefix length (%d)", cfg.StartPrefix, cfg.PrefixLength)
	}

}