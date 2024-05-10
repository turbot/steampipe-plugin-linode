package linode

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_token",
		Description: "Tokens in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listToken,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getToken,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "label", Type: proto.ColumnType_STRING, Description: "This token's label. This is for display purposes only, but can be used to more easily track what you're using each token for."},
			{Name: "token", Type: proto.ColumnType_STRING, Description: "First 16 characters of the token."},
			{Name: "scopes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Scopes").Transform(scopesStringToArray), Description: "Array of scopes for the token, e.g. *, account:read_write, domains:read_only."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time this token was created."},
			{Name: "expiry", Type: proto.ColumnType_TIMESTAMP, Description: "When this token will expire."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "This token's unique ID, which can be used to revoke it."},
		}),
	}
}

func listToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_token.listToken", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListTokens(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_token.listToken", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_token.getToken", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetToken(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_token.getToken", "query_error", err)
		return nil, err
	}
	return item, err
}

func scopesStringToArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return []string{}, nil
	}
	scopes := d.Value.(string)
	return strings.Split(scopes, " "), nil
}
