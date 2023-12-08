---
title: "Steampipe Table: linode_domain_record - Query Linode Domain Records using SQL"
description: "Allows users to query Linode Domain Records, providing detailed information about the DNS records associated with a domain."
---

# Table: linode_domain_record - Query Linode Domain Records using SQL

A Linode Domain Record represents a DNS record associated with a domain. These records define how internet traffic is directed for a domain. They can be created, updated, or deleted through the Linode API.

## Table Usage Guide

The `linode_domain_record` table provides insights into Domain Records within Linode. As a Network Administrator, explore record-specific details through this table, including record type, name, and associated data. Utilize it to uncover information about records, such as those associated with a specific domain, the record's target, and the priority of MX records.

**Important Notes**
- You must specify the `domain_id` in the `where` clause to query this table.

## Examples

### List all records for all domains
Explore the relationships between different domains and their associated records. This can help you understand the various connections and dependencies within your network infrastructure, providing valuable insights for management and troubleshooting purposes.

```sql+postgres
select
  d.domain,
  dr.record_type,
  dr.name,
  dr.target
from
  linode_domain as d,
  linode_domain_record as dr
where
  dr.domain_id = d.id;
```

```sql+sqlite
select
  d.domain,
  dr.record_type,
  dr.name,
  dr.target
from
  linode_domain as d,
  linode_domain_record as dr
where
  dr.domain_id = d.id;
```

### List all domain records for a domain
Explore all the domain records associated with a specific domain ID to understand its configuration and settings. This can be useful in managing and troubleshooting domain-related issues.

```sql+postgres
select
  *
from
  linode_domain_record
where
  domain_id = 1234;
```

```sql+sqlite
select
  *
from
  linode_domain_record
where
  domain_id = 1234;
```