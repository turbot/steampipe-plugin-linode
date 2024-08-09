package linode

import (
	"context"

	"github.com/linode/linodego"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinodeKubernetesCluster(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linode_kubernetes_cluster",
		Description: "Kubernetes clusters in the Linode account.",
		List: &plugin.ListConfig{
			Hydrate: listKubernetesCluster,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getKubernetesCluster,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "This Kubernetes cluster’s unique ID."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "This Kubernetes cluster’s unique label for display purposes only."},
			// Other columns
			{Name: "api_endpoints", Type: proto.ColumnType_JSON, Hydrate: listKubernetesClusterAPIEndpoints, Transform: transform.FromValue(), Description: "API endpoints for the cluster."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "When this Kubernetes cluster was created."},
			{Name: "k8s_version", Type: proto.ColumnType_STRING, Transform: transform.FromField("K8sVersion"), Description: "The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>, and the latest supported patch version will be deployed."},
			{Name: "kubeconfig", Type: proto.ColumnType_STRING, Hydrate: getKubeConfig, Transform: transform.FromField("KubeConfig").Transform(base64DecodedData), Description: "Kube config for the cluster."},
			{Name: "pools", Type: proto.ColumnType_JSON, Hydrate: listKubernetesClusterPools, Transform: transform.FromValue(), Description: "Pools for the cluster."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "This Kubernetes cluster’s location."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(transform.StringArrayToMap), Description: "Tags applied to the Kubernetes cluster as a map."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "List of Tags applied to the Kubernetes cluster."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "When this Kubernetes cluster was updated."},
		}),
	}
}

func listKubernetesCluster(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesCluster", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListLKEClusters(ctx, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesCluster", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getKubernetesCluster(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.getKubernetesCluster", "connection_error", err)
		return nil, err
	}
	item, err := conn.GetLKECluster(ctx, int(d.EqualsQuals["id"].GetInt64Value()))
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.getKubernetesCluster", "query_error", err)
		return nil, err
	}
	return item, err
}

func getKubeConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cluster := h.Item.(linodego.LKECluster)
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.getKubeConfig", "connection_error", err)
		return nil, err
	}
	item, err := conn.GetLKEClusterKubeconfig(ctx, cluster.ID)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.KubeConfig", "query_error", err, "cluster", cluster)
		return nil, err
	}
	return item, err
}

func listKubernetesClusterAPIEndpoints(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cluster := h.Item.(linodego.LKECluster)
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesClusterAPIEndpoints", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListLKEClusterAPIEndpoints(ctx, cluster.ID, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesClusterAPIEndpoints", "query_error", err, "cluster", cluster)
		return nil, err
	}
	endpoints := []string{}
	for _, i := range items {
		endpoints = append(endpoints, i.Endpoint)
	}
	return endpoints, err
}

func listKubernetesClusterPools(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cluster := h.Item.(linodego.LKECluster)
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesClusterPools", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListLKENodePools(ctx, cluster.ID, nil)
	if err != nil {
		plugin.Logger(ctx).Error("linode_kubernetes_cluster.listKubernetesClusterPools", "query_error", err, "cluster", cluster)
		return nil, err
	}
	return items, err
}
