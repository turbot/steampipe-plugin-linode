package linode

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type linodeConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &linodeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) linodeConfig {
	if connection == nil || connection.GetConfig() == nil {
		return linodeConfig{}
	}
	config, _ := connection.GetConfig().(linodeConfig)
	return config
}
