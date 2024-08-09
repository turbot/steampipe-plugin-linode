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

func tableLinodeEvent(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_event",
		Description: "Events in the Linode account.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "action", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
			},
			Hydrate: listEvent,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEvent,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique ID of this Event."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time this event was created."},
			{Name: "action", Type: proto.ColumnType_STRING, Description: "The action that caused this Event. New actions may be added in the future."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the Event, Enum: 'failed' 'finished' 'notification' 'scheduled' 'started'."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "The username of the User who caused the Event."},
			// Other columns
			{Name: "entity", Type: proto.ColumnType_JSON, Description: "Detailed information about the Event's entity, including ID, type, label, and URL used to access it."},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Raw Linode list filter string in JSON format."},
			{Name: "percent_complete", Type: proto.ColumnType_INT, Description: "A percentage estimating the amount of time remaining for an Event. Returns null for notification events."},
			{Name: "rate", Type: proto.ColumnType_STRING, Description: "The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events."},
			{Name: "read", Type: proto.ColumnType_BOOL, Description: "If this Event has been read."},
			{Name: "seen", Type: proto.ColumnType_BOOL, Description: "If this Event has been seen."},
			{Name: "secondary_entity", Type: proto.ColumnType_JSON, Description: "Detailed information about the Event's secondary or related entity, including ID, type, label, and URL used to access it."},
			{Name: "time_remaining", Type: proto.ColumnType_INT, Description: "The estimated time remaining until the completion of this Event. This value is only returned for in-progress events."},
		}),
	}
}

func listEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_event.listEvent", "connection_error", err)
		return nil, err
	}

	opts := linodego.ListOptions{}
	keyQuals := d.EqualsQuals
	if keyQuals["filter"] != nil {
		opts.Filter = keyQuals["filter"].GetStringValue()
	} else {
		filterParts := []string{}
		if keyQuals["action"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"action":"%s"`, keyQuals["action"].GetStringValue()))
		}
		if keyQuals["id"] != nil {
			filterParts = append(filterParts, fmt.Sprintf(`"id":%d`, keyQuals["id"].GetInt64Value()))
		}
		// TODO: created
		if len(filterParts) > 0 {
			opts.Filter = fmt.Sprintf("{%s}", strings.Join(filterParts, ","))
		}
	}

	items, err := conn.ListEvents(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("linode_event.listEvent", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_event.getEvent", "connection_error", err)
		return nil, err
	}
	id := int(d.EqualsQuals["id"].GetInt64Value())
	item, err := conn.GetEvent(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("linode_event.getEvent", "query_error", err)
		return nil, err
	}
	return item, err
}
