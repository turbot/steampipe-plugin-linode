package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeTag(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_tag",
		Description: "Tags in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listTag,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "label", Type: proto.ColumnType_STRING, Description: "A Label used for organization of objects on your Account."},
		}),
	}
}

func listTag(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_tag.listTag", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListTags(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_tag.listTag", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
