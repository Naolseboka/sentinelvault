package cfgvalid

import (
	"fmt"
	"net/url"
	"strings"
)

type ServiceConfig struct {
	Name string
	HTTPAddr string
	DatabaseURL string
	RedisURL string
	LogLevel string
}

func Validate(c ServiceConfig) []string {
	errs := make([]string, 0)
	if strings.TrimSpace(c.Name) == "" {
		errs = append(errs, "name required")
	}
	if c.HTTPAddr == "" {
		errs = append(errs, "http_addr required")
	}
	if c.DatabaseURL != "" {
		if _, err := url.Parse(c.DatabaseURL); err != nil {
			errs = append(errs, fmt.Sprintf("database_url: %v", err))
		}
	}
	if c.RedisURL != "" {
		if _, err := url.Parse(c.RedisURL); err != nil {
			errs = append(errs, fmt.Sprintf("redis_url: %v", err))
		}
	}
	switch strings.ToLower(c.LogLevel) {
	case "", "debug", "info", "warn", "error":
	default:
		errs = append(errs, "invalid log_level")
	}
	return errs
}

func MustValid(c ServiceConfig) error {
	if errs := Validate(c); len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	return nil
}
// maintenance note 91
