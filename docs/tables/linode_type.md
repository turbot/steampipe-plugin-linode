---
title: "Steampipe Table: linode_type - Query Linode Instance Types using SQL"
description: "Allows users to query Linode Instance Types, providing details about each instance type including its ID, label, disk, transfer, and more."
---

# Table: linode_type - Query Linode Instance Types using SQL

Linode Instance Types represent different hardware configurations that you can use for your Linode. Each type comes with different specifications, including CPU, memory, storage, and transfer capabilities. These types determine the capabilities and pricing of your Linode.

## Table Usage Guide

The `linode_type` table provides insights into the various instance types available within Linode. As a system administrator, you can use this table to understand each instance type's capabilities, including CPU, memory, storage, and transfer specifications. This information can be instrumental in choosing the most suitable instance type for your specific needs.

## Examples

### List all types
Analyze the settings to understand the different types of services available on Linode, sorted by their monthly cost. This is useful for budget planning and understanding the cost implications of different service types.

```sql+postgres
select
  *
from
  linode_type
order by
  price_monthly;
```

```sql+sqlite
select
  *
from
  linode_type
order by
  price_monthly;
```

### List all dedicated instance types
Explore the different dedicated instance types available on Linode, arranged in order of their monthly price. This can help you choose the most cost-effective option for your specific needs.

```sql+postgres
select
  *
from
  linode_type
where
  class = 'dedicated'
order by
  price_monthly;
```

```sql+sqlite
select
  *
from
  linode_type
where
  class = 'dedicated'
order by
  price_monthly;
```