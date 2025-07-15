package postgres_pool

import (
	"time"
)

type DSNConfig struct {
	Name            string
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

type DBConfigTemplate struct {
	Name         string                `mapstructure:"name"`
	Host         string                `mapstructure:"host"`
	Port         int                   `mapstructure:"port"`
	EnvVars      EnvVarMap             `mapstructure:"env_vars"`
	PoolSettings *PoolSettingsTemplate `mapstructure:"pool_settings"`
}

type PoolSettingsTemplate struct {
	MaxConnections        *int32  `mapstructure:"max_connections"`
	MinConnections        *int32  `mapstructure:"min_connections"`
	MaxConnectionLifetime *string `mapstructure:"max_connection_lifetime"`
	MinConnectionIdleTime *string `mapstructure:"min_connection_idle_time"`
}

type EnvVarMap struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type YamlConfig struct {
	Databases []DBConfigTemplate  `mapstructure:"databases"`
	Services  map[string][]string `mapstructure:"services"`
}

type ConfigTemplate struct {
	Databases []DBConfigTemplate  `mapstructure:"databases"`
	Services  map[string][]string `mapstructure:"services"`
}
