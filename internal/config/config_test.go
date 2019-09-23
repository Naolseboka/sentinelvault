package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFromYAMLAndEnv(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "svc.yaml")
	content := []byte(`
name: test-service
http_addr: ":8081"
database_url: postgres://localhost/test
feature_flags:
  beta: true
`)
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("FEATURE_ALPHA", "true")

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Name != "test-service" {
		t.Fatalf("name=%s", cfg.Name)
	}
	if cfg.LogLevel != "debug" {
		t.Fatalf("log level=%s", cfg.LogLevel)
	}
	if !cfg.FeatureEnabled("beta") || !cfg.FeatureEnabled("alpha") {
		t.Fatal("expected feature flags")
	}
}

func TestLoadMissingName(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.yaml")
	if err := os.WriteFile(path, []byte("http_addr: :8080"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Fatal("expected validation error")
	}
}
