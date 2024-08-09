package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinodeProfile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_profile",
		Description: "Profile of the user making the request.",
		List: &plugin.ListConfig{
			Hydrate: getProfile,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "uid", Type: proto.ColumnType_STRING, Description: "Your unique ID in our system. This value will never change, and can safely be used to identify your User."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Your username, used for logging in to our system."},
			// Other columns
			{Name: "authorized_keys", Type: proto.ColumnType_JSON, Description: "The list of SSH Keys authorized to use Lish for your User."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Your email address. This address will be used for communication with Linode as necessary."},
			{Name: "email_notifications", Type: proto.ColumnType_BOOL, Description: "If true, you will receive email notifications about account activity. If false, you may still receive business-critical communications through email."},
			{Name: "ip_whitelist_enabled", Type: proto.ColumnType_BOOL, Description: "If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled."},
			{Name: "lish_auth_method", Type: proto.ColumnType_STRING, Description: "The authentication methods that are allowed when connecting to the Linode Shell (Lish): password_keys, keys_only, disabled."},
			{Name: "referrals", Type: proto.ColumnType_JSON, Description: "Information about your status in our referral program. This information becomes accessible after this Profile’s Account has established at least $25.00 USD of total payments."},
			{Name: "restricted", Type: proto.ColumnType_BOOL, Description: "If true, your User has restrictions on what can be accessed on your Account."},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "The timezone you prefer to see times in."},
			{Name: "two_factor_auth", Type: proto.ColumnType_BOOL, Description: "If true, logins from untrusted computers will require Two Factor Authentication."},
		}),
	}
}

func getProfile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_profile.getProfile", "connection_error", err)
		return nil, err
	}
	item, err := conn.GetProfile(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("linode_profile.getProfile", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, item)
	return nil, nil
}
