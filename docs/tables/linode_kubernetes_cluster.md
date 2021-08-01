# Table: linode_kubernetes_cluster

List kubernetes clusters for the Linode account.

## Examples

### All clusters

```sql
select
  *
from
  linode_kubernetes_cluster
```

### Get Kube Config for a cluster

```sql
select
  kubeconfig
from
  linode_kubernetes_cluster
where
  id = 1234
```

### Instance details for each node in the cluster

```sql
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
  i.id = (node->'instance_id')::int
```
