package postgres_pool

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const (
	name     = "config"
	confType = "yaml"
	path     = "."
)

func LoadForCurrentService(envAppServiceName string) ([]DSNConfig, error) {
	serviceName := os.Getenv(envAppServiceName)
	if serviceName == "" {
		return nil, fmt.Errorf(
			"[ postgres_pool.LoadForCurrentService ] ERROR: environment variable %s is not set",
			envAppServiceName)
	}

	viper.SetConfigName(name)
	viper.SetConfigType(confType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("[ postgres_pool.LoadForCurrentService ] ERROR: error reading config file: %w", err)
	}

	var tpl ConfigTemplate
	if err := viper.Unmarshal(&tpl); err != nil {
		return nil, fmt.Errorf("[ postgres_pool.LoadForCurrentService ] ERROR: unable to decode config into struct: %w", err)
	}

	requiredDBNames, ok := tpl.Services[serviceName]
	if !ok {
		return nil, fmt.Errorf("[ postgres_pool.LoadForCurrentService ] ERROR: service with name '%s' not found in 'services' section, path %s", serviceName, path+name+"."+confType)
	}

	availableDBs := make(map[string]DBConfigTemplate)
	for _, dbTpl := range tpl.Databases {
		availableDBs[dbTpl.Name] = dbTpl
	}

	var finalConfigs []DSNConfig
	for _, requiredName := range requiredDBNames {
		dbTemplate, ok := availableDBs[requiredName]
		if !ok {
			return nil, fmt.Errorf("[ postgres_pool.LoadForCurrentService ] ERROR: database definition for '%s' (required by service '%s') not found in 'databases' section",
				requiredName, serviceName)
		}

		finalConfig, err := buildFinalConfig(dbTemplate)
		if err != nil {
			return nil, err
		}

		finalConfigs = append(finalConfigs, *finalConfig)
	}

	return finalConfigs, nil
}
