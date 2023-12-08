---
title: "Steampipe Table: linode_volume - Query Linode Volumes using SQL"
description: "Allows users to query Linode Volumes, providing insights into the block storage volumes associated with Linode instances."
---

# Table: linode_volume - Query Linode Volumes using SQL

Linode Volumes are a type of block storage that can be attached to Linode instances. They are highly available, automatically replicated, and allow users to create, resize, and move volumes within a Linode instance. Volumes are scalable and designed for 99.9% durability.

## Table Usage Guide

The `linode_volume` table provides insights into the block storage volumes within Linode. As a system administrator or DevOps engineer, you can explore volume-specific details through this table, including size, region, and associated Linode instance. Utilize it to uncover information about volumes, such as their current status, creation time, and the Linode instance they are attached to.

## Examples

### List volumes
Explore all available storage volumes within your Linode service. This can help you manage your storage resources effectively and plan for capacity upgrades if necessary.

```sql+postgres
select
  *
from
  linode_volume;
```

```sql+sqlite
select
  *
from
  linode_volume;
```

### Find volumes in a bad state
Discover the segments that have volumes in a suboptimal condition, requiring support intervention, to understand potential issues and regions affected. This aids in troubleshooting and maintaining system health.

```sql+postgres
select
  label,
  size,
  status,
  region
from
  linode_volume
where
  status = 'contact_support';
```

```sql+sqlite
select
  label,
  size,
  status,
  region
from
  linode_volume
where
  status = 'contact_support';
```

### Volumes with a given tag
Explore which storage volumes are associated with a specific tag to manage resources effectively across different regions.

```sql+postgres
select
  label,
  size,
  status,
  region
from
  linode_volume
where
  tags ? 'foo'
```

```sql+sqlite
Error: SQLite does not support the '?' operator for checking the existence of a key in a JSON object.
```

### Top 5 volumes by size
Discover the segments that have the largest volume sizes to better manage storage resources and optimize data distribution across different regions. This is particularly useful for identifying potential areas for data cleanup or redistribution to improve system performance.

```sql+postgres
select
  label,
  size,
  status,
  region
from
  linode_volume
order by
  size desc
limit
  5;
```

```sql+sqlite
select
  label,
  size,
  status,
  region
from
  linode_volume
order by
  size desc
limit
  5;
```