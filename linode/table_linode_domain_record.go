package linode

import (
	"context"

	"github.com/linode/linodego"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeDomainRecord(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_domain_record",
		Description: "Domain records for a given domain.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain_id"},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listDomainRecord,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"domain_id", "id"}),
			Hydrate:    getDomainRecord,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "domain_id", Type: proto.ColumnType_INT, Transform: transform.FromQual("domain_id"), Description: "The ID of the Domain for the record."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "This Record’s unique ID."},
			{Name: "record_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "The type of Record this is in the DNS system: A, AAAA, NS, MX, CNAME, TXT, SRV, PTR, CAA."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of this Record. This property’s actual usage and whether it is required depends on the type of record it represents. For example, for CNAME, it is the hostname."},
			// Other columns
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
			{Name: "port", Type: proto.ColumnType_INT, Description: "The port this Record points to. Only valid and required for SRV record requests."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Priority"), Description: "The priority of the target host for this Record. Lower values are preferred. Only valid for MX and SRV record requests. Required for SRV record requests."},
			{Name: "protocol", Type: proto.ColumnType_STRING, Description: "The protocol this Record’s service communicates with. An underscore (_) is prepended automatically to the submitted value for this property. Only valid for SRV record requests."},
			{Name: "service", Type: proto.ColumnType_STRING, Description: "The name of the service. Only valid and required for SRV record requests."},
			{Name: "tag", Type: proto.ColumnType_STRING, Description: "The tag portion of a CAA record. Only valid and required for CAA record requests."},
			{Name: "target", Type: proto.ColumnType_STRING, Description: "The target for this Record. For requests, this property’s actual usage and whether it is required depends on the type of record this represents. For example, for CNAME it is the domain target."},
			{Name: "ttl_sec", Type: proto.ColumnType_INT, Transform: transform.FromField("TTLSec").NullIfZero(), Description: "Time to Live - the amount of time in seconds that the domain record may be cached by resolvers or other domain servers."},
			{Name: "weight", Type: proto.ColumnType_INT, Description: "The relative weight of this Record used in the case of identical priority. Higher values are preferred. Only valid and required for SRV record requests."},
		}),
	}
}

func listDomainRecord(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain_record.listDomainRecord", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals

	domainID := int(keyQuals["domain_id"].GetInt64Value())

	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	}

	var items []linodego.DomainRecord
	if opts.Filter == "" {
		// An empty options object causes the SDK to panic?
		items, err = conn.ListDomainRecords(ctx, domainID, nil)
	} else {
		items, err = conn.ListDomainRecords(ctx, domainID, &opts)
	}
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain_record.listDomainRecord", "query_error", err, "domainID", domainID, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDomainRecord(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain_record.getDomainRecord", "connection_error", err)
		return nil, err
	}
	domainID := int(d.EqualsQuals["domain_id"].GetInt64Value())
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetDomainRecord(ctx, domainID, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_domain_record.getDomainRecord", "query_error", err)
		return nil, err
	}
	return item, err
}
