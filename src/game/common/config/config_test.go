package config

import (
	"testing"
)

type ServerConfig struct {
	Addr string
	Port int
}

type DbConfig struct {
	Url      string
	User     string
	Password string
}

type TestConfig struct {
	Server   ServerConfig
	Database DbConfig
}

// example for load json configure
func TestLoadJsonConfig(t *testing.T) {
	var config TestConfig
	err := LoadConfig("json", "example/json_example.json", &config)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("server={addr=%s, port=%d}", config.Server.Addr, config.Server.Port)
	t.Logf("database={url=%s, user=%s, password=%s}", config.Database.Url, config.Database.User, config.Database.Password)
}

func TestLoadXmlConfig(t *testing.T) {
	var config TestConfig
	err := LoadConfig("xml", "example/xml_example.xml", &config)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("addr=%s, port=%d", config.Server.Addr, config.Server.Port)
	t.Logf("database={url=%s, user=%s, password=%s}", config.Database.Url, config.Database.User, config.Database.Password)
}
