---
title: "Steampipe Table: linode_user - Query Linode Users using SQL"
description: "Allows users to query Linode Users, providing insights into user information, including their usernames, emails, and restricted status."
---

# Table: linode_user - Query Linode Users using SQL

Linode Users represent individual users within a Linode account. Each user can have different permissions and access levels, depending on their role within the account. Users can be granted either unrestricted access to all Linode resources or restricted access to only certain resources.

## Table Usage Guide

The `linode_user` table provides insights into individual users within a Linode account. As a system administrator, explore user-specific details through this table, including their usernames, emails, and whether they have restricted access. Utilize it to manage user permissions and access levels, ensuring the security and integrity of your Linode resources.

## Examples

### List users
Explore the list of users to understand who has access to your system, a crucial step in ensuring data security and managing user permissions.

```sql+postgres
select
  *
from
  linode_user;
```

```sql+sqlite
select
  *
from
  linode_user;
```