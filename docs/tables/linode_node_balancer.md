---
title: "Steampipe Table: linode_node_balancer - Query Linode NodeBalancers using SQL"
description: "Allows users to query Linode NodeBalancers, providing detailed information about each NodeBalancer's configurations, status, and specifications."
---

# Table: linode_node_balancer - Query Linode NodeBalancers using SQL

Linode NodeBalancers distribute incoming traffic across multiple Linode instances to ensure high availability and reliability. They are essential for managing and balancing the load on your Linode instances, allowing you to scale applications and services efficiently.

## Table Usage Guide

The `linode_node_balancer` table provides insights into each NodeBalancer within the Linode platform. As a system administrator or DevOps engineer, you can explore NodeBalancer-specific details through this table, including configurations, tags, and traffic metrics. Utilize it to gather information about each NodeBalancer, such as its current configuration, traffic handling, and region.

## Examples

### Basic info
Retrieve a list of all NodeBalancers in your Linode account to get an overview of your load balancing resources.

```sql+postgres
select
  id,
  created,
  updated,
  label,
  hostname
from
  linode_node_balancer;
```

```sql+sqlite
select
  id,
  created,
  updated,
  label,
  hostname
from
  linode_node_balancer;
```

### NodeBalancers by region
Explore which regions have the most NodeBalancers to better allocate resources and manage load distribution.

```sql+postgres
select
  region,
  count(*)
from
  linode_node_balancer
group by
  region;
```

```sql+sqlite
select
  region,
  count(*)
from
  linode_node_balancer
group by
  region;
```

### NodeBalancers created in the last 30 days
List all NodeBalancers that were created in the last 30 days to monitor new infrastructure additions.

```sql+postgres
select
  label,
  created,
  region
from
  linode_node_balancer
where
  created >= current_date - interval '30 days';
```

```sql+sqlite
select
  label,
  created,
  region
from
  linode_node_balancer
where
  created >= date('now', '-30 days');
```

### NodeBalancers by transfer usage
Identify NodeBalancers based on their transfer usage to manage bandwidth and optimize costs.

```sql+postgres
select
  label,
  transfer ->> 'in' as transfer_in,
  transfer ->> 'out' as transfer_out
from
  linode_node_balancer
order by
  (transfer ->> 'in')::bigint + (transfer ->> 'out')::bigint desc;
```

```sql+sqlite
select
  label,
  json_extract(transfer, '$.in') as transfer_in,
  json_extract(transfer, '$.out') as transfer_out
from
  linode_node_balancer
order by
  (json_extract(transfer, '$.in')) + (json_extract(transfer, '$.out')) desc;
```
