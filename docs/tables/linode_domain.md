---
title: "Steampipe Table: linode_domain - Query Linode Domains using SQL"
description: "Allows users to query Domains in Linode, specifically the domain and subdomain records, providing insights into DNS configurations and potential issues."
---

# Table: linode_domain - Query Linode Domains using SQL

Linode Domains is a service within Linode that allows you to manage and configure domain and subdomain records for your websites. It provides a centralized way to set up and manage DNS records for various Linode resources, including virtual machines, databases, web applications, and more. Linode Domains helps you stay informed about the health and performance of your DNS configurations and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `linode_domain` table provides insights into Domains within Linode. As a DevOps engineer, explore domain-specific details through this table, including domain and subdomain records, TTL values, and associated metadata. Utilize it to uncover information about domains, such as those with misconfigured DNS records, the relationships between domains and subdomains, and the verification of TTL values.

## Examples

### List all domains
Explore all the domains available in your Linode account. This can help in managing and organizing your resources effectively.

```sql+postgres
select
  *
from
  linode_domain;
```

```sql+sqlite
select
  *
from
  linode_domain;
```

### Domains with a given tag
Discover the segments that are tagged with a specific label, enabling you to organize and manage your resources more effectively. This is particularly useful when you need to perform actions on a group of resources that share a common tag.

```sql+postgres
select
  domain,
  tags
from
  linode_volume
where
  tags ? 'foo'
```

```sql+sqlite
Error: SQLite does not support '?' operator for checking the existence of a key in a JSON object.
```