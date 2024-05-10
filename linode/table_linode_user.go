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

func tableLinodeUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_user",
		Description: "Users in the Linode account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "username", Require: plugin.Optional},
			},
			Hydrate: listUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("username"),
			Hydrate:    getUser,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "username", Type: proto.ColumnType_STRING, Description: "This User’s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts)."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address for this User, for account management communications, and may be used for other communications as configured."},
			{Name: "restricted", Type: proto.ColumnType_BOOL, Description: "If true, this User must be granted access to perform actions or access entities on this Account."},
			{Name: "ssh_keys", Type: proto.ColumnType_JSON, Description: "A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request."},
			// Other columns
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
		}),
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_user.listUser", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		if keyQuals["username"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"username":"%s"`, keyQuals["username"].GetStringValue()))
		}
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListUsers(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_user.listUser", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_user.getUser", "connection_error", err)
		return nil, err
	}
	username := d.EqualsQuals["username"].GetStringValue()
	item, err := conn.GetUser(ctx, username)
	if err != nil {
		plugin.Logger(ctx).Error("linode_user.getUser", "query_error", err)
		return nil, err
	}
	return item, err
}
