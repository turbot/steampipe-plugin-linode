---
title: "Steampipe Table: linode_tag - Query Linode Tags using SQL"
description: "Allows users to query Linode Tags, specifically the tag names and entities associated with them, providing insights into resource categorization and organization."
---

# Table: linode_tag - Query Linode Tags using SQL

Linode Tags are a feature in Linode that allows users to organize and categorize their Linode resources. These tags can be applied to various resources including Linodes, Block Storage Volumes, NodeBalancers, Domains, and more. They provide an efficient way to manage resources, especially in larger environments with multiple resources.

## Table Usage Guide

The `linode_tag` table provides insights into the tags used within Linode. As a system administrator or DevOps engineer, explore tag-specific details through this table, including associated resources and their types. Utilize it to uncover information about resource organization, such as which resources are associated with certain tags, aiding in efficient resource management and categorization.

## Examples

### List all tags
Explore all the tags available in your Linode account, sorted by their labels. This is useful for understanding how resources are categorized and can aid in managing and organizing your resources effectively.

```sql+postgres
select
  *
from
  linode_tag
order by
  label;
```

```sql+sqlite
select
  *
from
  linode_tag
order by
  label;
```