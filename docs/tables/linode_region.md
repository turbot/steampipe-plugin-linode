---
title: "Steampipe Table: linode_region - Query Linode Regions using SQL"
description: "Allows users to query Linode Regions, specifically to retrieve information about the various regions where Linode services are available."
---

# Table: linode_region - Query Linode Regions using SQL

A Linode Region represents a geographic location where Linode services are hosted. These regions are data centers that are spread across the world, and each region may have multiple availability zones. Choosing the right region can help reduce latency and increase redundancy for your applications.

## Table Usage Guide

The `linode_region` table provides insights into the geographical distribution of Linode services. As a system administrator or DevOps engineer, you can use this table to explore details about each region, including its location and capabilities. This information can be crucial when planning the deployment of applications, ensuring optimal performance and redundancy.

## Examples

### List all regions
Discover the segments that allow you to understand the geographical distribution of your Linode resources, helping you manage and optimize your operations based on regional factors.

```sql+postgres
select
  *
from
  linode_region
order by
  id;
```

```sql+sqlite
select
  *
from
  linode_region
order by
  id;
```

### List all US regions
Explore all regions within the United States to better understand the geographical distribution of your resources. This can assist in optimizing resource allocation and improving operational efficiency.

```sql+postgres
select
  *
from
  linode_region
where
  country = 'us'
order by
  id;
```

```sql+sqlite
select
  *
from
  linode_region
where
  country = 'us'
order by
  id;
```