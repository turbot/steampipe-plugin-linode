package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-linode",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"linode_account":            tableLinodeAccount(ctx),
			"linode_bucket":             tableLinodeBucket(ctx),
			"linode_domain":             tableLinodeDomain(ctx),
			"linode_domain_record":      tableLinodeDomainRecord(ctx),
			"linode_event":              tableLinodeEvent(ctx),
			"linode_firewall":           tableLinodeFirewall(ctx),
			"linode_image":              tableLinodeImage(ctx),
			"linode_instance":           tableLinodeInstance(ctx),
			"linode_kubernetes_cluster": tableLinodeKubernetesCluster(ctx),
			"linode_node_balancer":      tableLinodeNodeBalancer(ctx),
			"linode_profile":            tableLinodeProfile(ctx),
			"linode_region":             tableLinodeRegion(ctx),
			"linode_tag":                tableLinodeTag(ctx),
			"linode_token":              tableLinodeToken(ctx),
			"linode_type":               tableLinodeType(ctx),
			"linode_user":               tableLinodeUser(ctx),
			"linode_volume":             tableLinodeVolume(ctx),
		},
	}
	return p
}
