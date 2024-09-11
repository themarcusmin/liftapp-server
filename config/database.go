package config

import (
	"time"
)

// DatabaseConfig - all database variables
type DatabaseConfig struct {
	// relational database
	RDBMS RDBMS
}

// RDBMS - relational database variables
type RDBMS struct {
	Activate string
	Env      struct {
		Driver   string
		Host     string
		Port     string
		TimeZone string
	}
	Access struct {
		DbName string
		User   string
		Pass   string
	}
	Ssl struct {
		Sslmode    string
		MinTLS     string
		RootCA     string
		ServerCert string
		ClientCert string
		ClientKey  string
	}
	Conn struct {
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
	}
	Log struct {
		LogLevel int
	}
}
