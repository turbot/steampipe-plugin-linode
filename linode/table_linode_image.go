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

func tableLinodeImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_image",
		Description: "Images in the Linode account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "deprecated", Require: plugin.Optional},
				{Name: "is_public", Require: plugin.Optional},
				{Name: "label", Require: plugin.Optional},
				{Name: "size", Require: plugin.Optional},
				{Name: "image_type", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getImage,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of this Image."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "A short description of the Image."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When this Image was created."},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: "The name of the User who created this Image, or 'linode' for official Images."},
			{Name: "deprecated", Type: proto.ColumnType_BOOL, Description: "Whether or not this Image is deprecated. Will only be true for deprecated public Images."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A detailed description of this Image."},
			{Name: "expiry", Type: proto.ColumnType_TIMESTAMP, Description: "Only Images created automatically (from a deleted Linode; type=automatic) will expire."},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
			{Name: "image_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("type"), Description: "How the Image was created: manual, automatic."},
			{Name: "is_public", Type: proto.ColumnType_BOOL, Description: "True if the Image is public."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "The minimum size this Image needs to deploy. Size is in MB."},
			{Name: "vendor", Type: proto.ColumnType_STRING, Description: "The upstream distribution vendor. None for private Images."},
		}),
	}
}

func listImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_image.listImage", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		if keyQuals["size"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"size":%d`, int(keyQuals["size"].GetInt64Value())))
		}
		if keyQuals["deprecated"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"deprecated":%t`, keyQuals["deprecated"].GetBoolValue()))
		}
		if keyQuals["is_public"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"is_public":%t`, keyQuals["is_public"].GetBoolValue()))
		}
		if keyQuals["image_type"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"type":"%s"`, keyQuals["image_type"].GetStringValue()))
		}
		if keyQuals["label"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"label":"%s"`, keyQuals["label"].GetStringValue()))
		}
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListImages(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_image.listImage", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_image.getImage", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	item, err := conn.GetImage(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_image.getImage", "query_error", err)
		return nil, err
	}
	return item, err
}
