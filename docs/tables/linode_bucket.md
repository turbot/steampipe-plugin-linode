---
title: "Steampipe Table: linode_bucket - Query Linode Object Storage Buckets using SQL"
description: "Allows users to query Linode Object Storage Buckets, providing insights into the metadata and configuration settings of each bucket."
---

# Table: linode_bucket - Query Linode Object Storage Buckets using SQL

Linode Object Storage is a globally-available, S3-compatible method for storing and accessing data. It offers highly scalable, robust, and inexpensive storage for backup, archiving, content distribution, and more. With Linode Object Storage, your data is stored and served in a secure and distributed manner.

## Table Usage Guide

The `linode_bucket` table provides insights into Object Storage Buckets within Linode's Object Storage service. As a DevOps engineer, explore bucket-specific details through this table, including names, created timestamps, and cluster regions. Utilize it to uncover information about buckets, such as their size, the number of objects they contain, and their respective cluster regions.

## Examples

### List buckets
Explore all the storage buckets available in your Linode account. This can help manage resources and assess storage needs.

```sql+postgres
select
  *
from
  linode_bucket;
```

```sql+sqlite
select
  *
from
  linode_bucket;
```

### Buckets in us-east-1
Explore which Linode storage buckets are located in the 'us-east-1' region. This could be useful for managing data locality and ensuring regulatory compliance.

```sql+postgres
select
  *
from
  linode_bucket
where
  cluster = 'us-east-1';
```

```sql+sqlite
select
  *
from
  linode_bucket
where
  cluster = 'us-east-1';
```