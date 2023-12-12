package linode

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type linodeConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &linodeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) linodeConfig {
	if connection == nil || connection.Config == nil {
		return linodeConfig{}
	}
	config, _ := connection.Config.(linodeConfig)
	return config
}
