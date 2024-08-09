package linode

import (
	"context"

	"github.com/linode/linodego"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeNodeBalancer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_node_balancer",
		Description: "NodeBalancers that are assigned to this Linode and readable by the requesting User.",
		List: &plugin.ListConfig{
			Hydrate: listNodeBalancers,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique ID of this NodeBalancer."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When the NodeBalancer was created."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "When the NodeBalancer was updated."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The NodeBalancer's label. These must be unique on your Account."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The Region where this NodeBalancer is located. NodeBalancers only support backends in the same Region."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The NodeBalancer's hostname, ending with .nodebalancer.linode.com."},
			{Name: "ipv4", Type: proto.ColumnType_IPADDR, Description: "The NodeBalancer's public IPv4 address.", Transform: transform.FromField("IPv4")},
			{Name: "ipv6", Type: proto.ColumnType_IPADDR, Description: "The NodeBalancer's public IPv6 address.", Transform: transform.FromField("IPv6")},
			{Name: "client_conn_throttle", Type: proto.ColumnType_INT, Description: "Throttle connections per second (0-20). Set to 0 (zero) to disable throttling."},
			{Name: "transfer", Type: proto.ColumnType_JSON, Description: "Information about the amount of transfer this NodeBalancer has had so far this month."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "An array of tags applied to this object. Tags are for organizational purposes only."},
			{Name: "configurations", Hydrate: getNodeBalancersConfiguration, Transform: transform.FromValue(), Type: proto.ColumnType_JSON, Description: "The NodeBalancer configurations."},
		}),
	}
}

//// LIST FUNCTION

func listNodeBalancers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_node_balancer.listNodeBalancers", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}

	items, err := conn.ListNodeBalancers(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_node_balancer.listNodeBalancers", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTION

func getNodeBalancersConfiguration(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	nodeBalanccer := h.Item.(linodego.NodeBalancer)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_node_balancer.getNodeBalancersConfiguration", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}

	items, err := conn.ListNodeBalancerConfigs(ctx, nodeBalanccer.ID, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_node_balancer.getNodeBalancersConfiguration", "query_error", err, "opts", opts)
		return nil, err
	}
	if len(items) > 0 {
		return items, nil
	}
	return nil, nil
}
