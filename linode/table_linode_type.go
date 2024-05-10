package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_type",
		Description: "Types in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listType,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getType,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID representing the Linode Type."},
			{Name: "disk", Type: proto.ColumnType_INT, Description: "The Disk size, in MB, of the Linode Type."},
			{Name: "class", Type: proto.ColumnType_STRING, Description: "The class of the Linode Type: nanode, standard, dedicated, gpu, highmem."},
			{Name: "price_hourly", Type: proto.ColumnType_INT, Transform: transform.FromField("Price.Hourly"), Description: "Cost (in US dollars) per hour."},
			{Name: "price_monthly", Type: proto.ColumnType_INT, Transform: transform.FromField("Price.Monthly"), Description: "Cost (in US dollars) per month."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The Linode Type’s label is for display purposes only."},
			{Name: "addons", Type: proto.ColumnType_JSON, Description: "A list of optional add-on services for Linodes and their associated costs."},
			{Name: "network_out", Type: proto.ColumnType_INT, Description: "The Mbits outbound bandwidth allocation."},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "Amount of RAM included in this Linode Type."},
			{Name: "transfer", Type: proto.ColumnType_INT, Description: "The monthly outbound transfer amount, in MB."},
			{Name: "vcpus", Type: proto.ColumnType_INT, Transform: transform.FromField("VCPUs"), Description: "The number of VCPU cores this Linode Type offers."},
		}),
	}
}

func listType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_type.listType", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListTypes(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_type.listType", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_type.getType", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	item, err := conn.GetType(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_type.getType", "query_error", err)
		return nil, err
	}
	return item, err
}
