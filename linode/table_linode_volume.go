package linode

import (
	"context"
	"fmt"
	"strings"

	"github.com/linode/linodego"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_volume",
		Description: "Volumes in the Linode account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "label", Require: plugin.Optional},
				//{Name: "tags", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolume,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique ID of this Volume."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The Volume’s label is for display purposes only."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "The Volume’s size, in GiB."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When this Volume was created."},
			{Name: "filesystem_path", Type: proto.ColumnType_STRING, Description: "The full filesystem path for the Volume based on the Volume’s label."},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
			{Name: "linode_id", Type: proto.ColumnType_INT, Description: "If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Region where the volume resides."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the volume: creating, active, resizing, contact_support."},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(transform.StringArrayToMap), Description: "Tags applied to this volume as a map."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "List of Tags applied to this volume."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "When this Volume was last updated."},
		}),
	}
}

func listVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_volume.listVolume", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		if keyQuals["label"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"label":"%s"`, keyQuals["label"].GetStringValue()))
		}
		// TODO: tags
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListVolumes(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_volume.listVolume", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_volume.getVolume", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetVolume(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_volume.getVolume", "query_error", err)
		return nil, err
	}
	return item, err
}
