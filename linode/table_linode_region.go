package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_region",
		Description: "Regions in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listRegion,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRegion,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The region."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country for the region."},
		}),
	}
}

func listRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_region.listRegion", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListRegions(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_region.listRegion", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_region.getRegion", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	item, err := conn.GetRegion(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_region.getRegion", "query_error", err)
		return nil, err
	}
	return item, err
}
