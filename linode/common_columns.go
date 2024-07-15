package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "euuid",
			Description: "An external unique identifier for this account.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAccountEuuid,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getAccountEuuidMemoized = plugin.HydrateFunc(getAccountEuuidUncached).Memoize(memoize.WithCacheKeyFunction(getAccountEuuidCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccountEuuid(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountEuuidMemoized(ctx, d, h)
}

// Build a cache key for the call to getAccountEuuidCacheKey.
func getAccountEuuidCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getAccountEuuid"
	return key, nil
}

func getAccountEuuidUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_account_settings.getAccount", "connection_error", err)
		return nil, err
	}
	item, err := conn.GetAccount(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("linode_account_settings.getAccount", "query_error", err)
		return nil, err
	}

	return item.EUUID, nil
}