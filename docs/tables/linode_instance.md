---
title: "Steampipe Table: linode_instance - Query Linode Instances using SQL"
description: "Allows users to query Linode Instances, providing detailed information about each instance's configurations, status, and specifications."
---

# Table: linode_instance - Query Linode Instances using SQL

Linode Instances are virtual machines that run on the Linode platform. They are the building blocks of Linode and can be easily created, scaled, managed, and deleted. Each instance is a separate virtual machine within the same physical server.

## Table Usage Guide

The `linode_instance` table provides insights into each Linode Instance within the Linode platform. As a system administrator or a DevOps engineer, you can explore instance-specific details through this table, including configurations, status, and specifications. Utilize it to uncover information about each instance, such as its current running status, its hardware specifications, and its networking configurations.

## Examples

### List all instances
Explore all the instances within your Linode account to gain a comprehensive overview of your resources. This can help in efficient resource management and in identifying any anomalies or areas for optimization.

```sql+postgres
select
  *
from
  linode_instance;
```

```sql+sqlite
select
  *
from
  linode_instance;
```

### Instances by region
Explore which regions have the most instances to better allocate resources and manage server load.

```sql+postgres
select
  region,
  count(*)
from
  linode_instance
group by
  region;
```

```sql+sqlite
select
  region,
  count(*)
from
  linode_instance
group by
  region;
```

### Instances by type
Explore which types of instances are most commonly used, allowing you to understand usage patterns and optimize resource allocation.

```sql+postgres
select
  instance_type,
  count(*)
from
  linode_instance
group by
  instance_type;
```

```sql+sqlite
select
  instance_type,
  count(*)
from
  linode_instance
group by
  instance_type;
```

### Instances with a given tag
Explore which instances have been tagged with a specific label. This is useful for categorizing and quickly identifying instances in your Linode environment.

```sql+postgres
select
  label,
  status,
  tags
from
  linode_instance
where
  tags ? 'foo'
```

```sql+sqlite
Error: SQLite does not support '?' operator used for checking if an element exists in a JSON array.
```

### Running instances
Explore which Linode instances are currently running to manage resources effectively and ensure optimal performance. This is beneficial in real-world scenarios for maintaining server health and preventing overloads.

```sql+postgres
select
  label,
  ipv4,
  status,
  tags
from
  linode_instance
where
  status = 'running';
```

```sql+sqlite
select
  label,
  ipv4,
  status,
  tags
from
  linode_instance
where
  status = 'running';
```