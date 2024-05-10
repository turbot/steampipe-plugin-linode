package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeBucket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_bucket",
		Description: "Buckets in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listBucket,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"cluster", "label"}),
			Hydrate:    getBucket,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "label", Type: proto.ColumnType_STRING, Description: "The name of this bucket."},
			// Other columns
			{Name: "cluster", Type: proto.ColumnType_STRING, Description: "The ID of the Object Storage Cluster this bucket is in."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When this bucket was created."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public."},
		}),
	}
}

func listBucket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_bucket.listBucket", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListObjectStorageBuckets(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_bucket.listBucket", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getBucket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_bucket.getBucket", "connection_error", err)
		return nil, err
	}
	item, err := conn.GetObjectStorageBucket(ctx, d.EqualsQuals["cluster"].GetStringValue(), d.EqualsQuals["label"].GetStringValue())
	if err != nil {
		plugin.Logger(ctx).Error("linode_bucket.getBucket", "query_error", err)
		return nil, err
	}
	return item, err
}
