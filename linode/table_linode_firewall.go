package linode

import (
	"context"

	"github.com/linode/linodego"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeFirewall(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_firewall",
		Description: "Firewalls in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listFirewalls,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFirewall,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique ID of this Firewall."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time this firewall was created."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time this firewall was updated."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The firewall’s label is for display purposes only."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the firewall. Possible values are 'enabled', 'disabled', or 'deleted'."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags applied to this firewall."},
			{Name: "rules", Type: proto.ColumnType_JSON, Description: "The rules associated with the firewall."},
		}),
	}
}

func listFirewalls(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_firewall.listFirewall", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}

	items, err := conn.ListFirewalls(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_firewall.listFirewall", "query_error", err, "opts", opts)
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

func getFirewall(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_firewall.getFirewall", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetFirewall(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_firewall.getFirewall", "query_error", err)
		return nil, err
	}
	return item, err
}
