---
title: "Steampipe Table: linode_kubernetes_cluster - Query Linode Kubernetes Clusters using SQL"
description: "Allows users to query Linode Kubernetes Clusters, providing detailed information about each cluster's configuration, nodes, and status."
---

# Table: linode_kubernetes_cluster - Query Linode Kubernetes Clusters using SQL

Linode Kubernetes Engine (LKE) is a managed service within Linode that allows you to deploy, manage, and scale containerized applications using Kubernetes. It simplifies the process of managing clusters by handling the underlying infrastructure, allowing you to focus on deploying applications. LKE provides a highly available, scalable, and secure environment for running your applications.

## Table Usage Guide

The `linode_kubernetes_cluster` table provides insights into Kubernetes clusters within Linode Kubernetes Engine (LKE). As a DevOps engineer, explore cluster-specific details through this table, including the configuration, nodes, and status. Utilize it to uncover information about clusters, such as their current status, node count, and Kubernetes version.

## Examples

### All clusters
Explore all the Kubernetes clusters in your Linode account to manage and optimize your cloud resources effectively. This can help you in identifying underutilized resources and potential areas for cost savings.

```sql+postgres
select
  *
from
  linode_kubernetes_cluster;
```

```sql+sqlite
select
  *
from
  linode_kubernetes_cluster;
```

### Get Kube Config for a cluster
Review the configuration for a specific Kubernetes cluster within the Linode service. This is useful for understanding the cluster's setup and managing its resources.

```sql+postgres
select
  kubeconfig
from
  linode_kubernetes_cluster
where
  id = 1234;
```

```sql+sqlite
select
  kubeconfig
from
  linode_kubernetes_cluster
where
  id = 1234;
```

### Instance details for each node in the cluster
Explore the operational status and other relevant details of each node within a specific cluster. This can be useful in monitoring and managing resources within a cloud-based infrastructure.

```sql+postgres
select
  node ->> 'id' as node_id,
  node ->> 'status' as node_status,
  i.*
from
  linode_kubernetes_cluster as kc,
  jsonb_array_elements(kc.pools) as pool,
  jsonb_array_elements(pool->'nodes') as node,
  linode_instance as i
where
  i.id = (node->'instance_id')::int;
```

```sql+sqlite
select
  json_extract(node.value, '$.id') as node_id,
  json_extract(node.value, '$.status') as node_status,
  i.*
from
  linode_kubernetes_cluster as kc,
  json_each(kc.pools) as pool,
  json_each(json_extract(pool.value, '$.nodes')) as node,
  linode_instance as i
where
  i.id = json_extract(node.value, '$.instance_id');
```