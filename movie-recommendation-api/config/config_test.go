package config

import (
    "os"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("DB_HOST", "localhost")
    os.Setenv("DB_PORT", "5432")
    os.Setenv("DB_USER", "user")
    os.Setenv("DB_PASSWORD", "password")
    os.Setenv("DB_NAME", "testdb")
    os.Setenv("TMDB_API_KEY", "testapikey")

    cfg, err := LoadConfig()
    if err != nil {
        t.Errorf("LoadConfig() error = %v", err)
    }

    if cfg.DBHost != "localhost" || cfg.DBName != "testdb" {
        t.Errorf("LoadConfig() did not load env vars correctly")
    }

    expectedConnectionString := "host=localhost port=5432 user=user password=password dbname=testdb sslmode=disable"
    if cfg.DBConnectionString() != expectedConnectionString {
        t.Errorf("DBConnectionString() = %v, want %v", cfg.DBConnectionString(), expectedConnectionString)
    }
}

func TestDBConnectionString(t *testing.T) {
    cfg := &Config{
        DBHost:     "localhost",
        DBPort:     "5432",
        DBUser:     "user",
        DBPassword: "password",
        DBName:     "testdb",
    }

    expected := "host=localhost port=5432 user=user password=password dbname=testdb sslmode=disable"
    actual := cfg.DBConnectionString()
    if actual != expected {
        t.Errorf("DBConnectionString() incorrect, got %v, want %v", actual, expected)
    }
}
