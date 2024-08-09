package linode

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_account",
		Description: "Account information.",
		List: &plugin.ListConfig{
			Hydrate: getAccount,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address of the person associated with this Account."},
			// Other columns
			{Name: "address_1", Type: proto.ColumnType_STRING, Description: "First line of this Account’s billing address."},
			{Name: "address_2", Type: proto.ColumnType_STRING, Description: "Second line of this Account’s billing address."},
			{Name: "balance", Type: proto.ColumnType_STRING, Description: "This Account’s balance, in US dollars."},
			{Name: "balance_uninvoiced", Type: proto.ColumnType_STRING, Description: "This Account’s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate."},
			{Name: "city", Type: proto.ColumnType_STRING, Description: "The city for this Account’s billing address."},
			{Name: "company", Type: proto.ColumnType_STRING, Description: "The company name associated with this Account."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "The two-letter country code of this Account’s billing address."},
			{Name: "credit_card", Type: proto.ColumnType_JSON, Description: "Credit Card information associated with this Account."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "The first name of the person associated with this Account."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "The last name of the person associated with this Account."},
			{Name: "euuid", Type: proto.ColumnType_STRING, Description: "An external unique identifier for this account.", Transform: transform.FromField("EUUID")},
			{Name: "phone", Type: proto.ColumnType_STRING, Description: "The phone number associated with this Account."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The state for this Account’s billing address."},
			{Name: "tax_id", Type: proto.ColumnType_STRING, Description: "The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be null."},
			{Name: "zip", Type: proto.ColumnType_STRING, Description: "The zip code of this Account’s billing address."},
		},
	}
}

func getAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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
	d.StreamListItem(ctx, *item)
	return nil, nil
}
