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

func tableLinodeDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_domain",
		Description: "Domains in the Linode account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain", Require: plugin.Optional},
				//{Name: "tags", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDomain,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique ID of this Domain."},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain."},
			// Other columns
			{Name: "axfr_ips", Type: proto.ColumnType_JSON, Transform: transform.FromField("AXfrIPs"), Description: "The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description for this Domain. This is for display purposes only."},
			{Name: "domain_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave)."},
			{Name: "expire_sec", Type: proto.ColumnType_INT, Description: "The amount of time in seconds that may pass before this Domain is no longer authoritative."},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
			{Name: "master_ips", Type: proto.ColumnType_JSON, Transform: transform.FromField("MasterIPs"), Description: "The IP addresses representing the master DNS for this Domain."},
			{Name: "refresh_sec", Type: proto.ColumnType_INT, Description: "The amount of time in seconds before this Domain should be refreshed."},
			{Name: "retry_sec", Type: proto.ColumnType_INT, Description: "The interval, in seconds, at which a failed refresh should be retried."},
			{Name: "soa_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("SOAEmail"), Description: "Start of Authority email address. This is required for master Domains."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Used to control whether this Domain is currently being rendered: disabled, active, edit_mode, has_errors."},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(transform.StringArrayToMap), Description: "Tags applied to this domain as a map."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "List of Tags applied to this domain."},
			{Name: "ttl_sec", Type: proto.ColumnType_INT, Transform: transform.FromField("TTLSec").NullIfZero(), Description: "Time to Live - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers."},
		}),
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain.listDomain", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		if keyQuals["domain"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"domain":"%s"`, keyQuals["domain"].GetStringValue()))
		}
		// TODO: tags
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListDomains(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain.listDomain", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain.getDomain", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetDomain(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain.getDomain", "query_error", err)
		return nil, err
	}
	return item, err
}
