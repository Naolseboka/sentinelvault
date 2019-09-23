package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// ServiceConfig holds runtime configuration for a microservice.
type ServiceConfig struct {
	Name          string            `yaml:"name" json:"name"`
	HTTPAddr      string            `yaml:"http_addr" json:"http_addr"`
	GRPCAddr      string            `yaml:"grpc_addr" json:"grpc_addr"`
	DatabaseURL   string            `yaml:"database_url" json:"database_url"`
	RedisURL      string            `yaml:"redis_url" json:"redis_url"`
	KafkaBrokers  []string          `yaml:"kafka_brokers" json:"kafka_brokers"`
	LogLevel      string            `yaml:"log_level" json:"log_level"`
	OTelEndpoint  string            `yaml:"otel_endpoint" json:"otel_endpoint"`
	FeatureFlags  map[string]bool   `yaml:"feature_flags" json:"feature_flags"`
	Extra         map[string]string `yaml:"extra" json:"extra"`
}

// Load reads configuration from optional YAML path and environment overrides.
func Load(yamlPath string) (*ServiceConfig, error) {
	cfg := defaultConfig()
	if yamlPath != "" {
		data, err := os.ReadFile(yamlPath)
		if err != nil {
			return nil, fmt.Errorf("read config %s: %w", yamlPath, err)
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parse config %s: %w", yamlPath, err)
		}
	}
	applyEnv(cfg)
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func defaultConfig() *ServiceConfig {
	return &ServiceConfig{
		HTTPAddr:     ":8080",
		GRPCAddr:     ":9090",
		LogLevel:     "info",
		FeatureFlags: map[string]bool{},
		Extra:        map[string]string{},
	}
}

func applyEnv(cfg *ServiceConfig) {
	if v := os.Getenv("SERVICE_NAME"); v != "" {
		cfg.Name = v
	}
	if v := os.Getenv("HTTP_ADDR"); v != "" {
		cfg.HTTPAddr = v
	}
	if v := os.Getenv("GRPC_ADDR"); v != "" {
		cfg.GRPCAddr = v
	}
	if v := os.Getenv("DATABASE_URL"); v != "" {
		cfg.DatabaseURL = v
	}
	if v := os.Getenv("REDIS_URL"); v != "" {
		cfg.RedisURL = v
	}
	if v := os.Getenv("KAFKA_BROKERS"); v != "" {
		cfg.KafkaBrokers = splitCSV(v)
	}
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		cfg.LogLevel = v
	}
	if v := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"); v != "" {
		cfg.OTelEndpoint = v
	}
	for _, e := range os.Environ() {
		const prefix = "FEATURE_"
		if strings.HasPrefix(e, prefix) {
			parts := strings.SplitN(e, "=", 2)
			if len(parts) != 2 {
				continue
			}
			key := strings.ToLower(strings.TrimPrefix(parts[0], prefix))
			enabled, _ := strconv.ParseBool(parts[1])
			if cfg.FeatureFlags == nil {
				cfg.FeatureFlags = map[string]bool{}
			}
			cfg.FeatureFlags[key] = enabled
		}
	}
}

func (c *ServiceConfig) validate() error {
	if c.Name == "" {
		return fmt.Errorf("service name is required")
	}
	if c.HTTPAddr == "" {
		return fmt.Errorf("http_addr is required")
	}
	return nil
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// FeatureEnabled returns whether a feature flag is enabled.
func (c *ServiceConfig) FeatureEnabled(name string) bool {
	if c.FeatureFlags == nil {
		return false
	}
	return c.FeatureFlags[name]
}
