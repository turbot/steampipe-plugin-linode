package main

import (
	"github.com/turbot/steampipe-plugin-linode/linode"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linode.Plugin})
}
