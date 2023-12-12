---
title: "Steampipe Table: linode_token - Query Linode Personal Access Tokens using SQL"
description: "Allows users to query Linode Personal Access Tokens, providing insights into token details, including its scopes, created and expiry dates."
---

# Table: linode_token - Query Linode Personal Access Tokens using SQL

A Linode Personal Access Token (PAT) is an API token granting full access to your Linode account, with the ability to create, access, modify, and remove any aspect of your account. PATs are great for when you want to build your own applications or scripts. Just like OAuth tokens, you can limit a token's scope to restrict what parts of your account it can access.

## Table Usage Guide

The `linode_token` table provides insights into Personal Access Tokens within Linode. As a DevOps engineer, explore token-specific details through this table, including its scopes, created and expiry dates. Utilize it to uncover information about tokens, such as those with certain scopes, and verification of expiry dates.

## Examples

### List tokens
Explore all the available tokens in your Linode account to manage and control access to your resources effectively. This helps in maintaining security and managing permissions within your account.

```sql+postgres
select
  *
from
  linode_token;
```

```sql+sqlite
select
  *
from
  linode_token;
```

### Tokens by age in days
Identify the age of your Linode tokens to understand how long they have been in existence. This can help in managing token lifecycle and ensuring security compliance by replacing or revoking old tokens.

```sql+postgres
select
  token,
  created,
  date_part('day', age(current_timestamp, created)) as age_days
from
  linode_token
order by
  age_days desc;
```

```sql+sqlite
select
  token,
  created,
  julianday('now') - julianday(created) as age_days
from
  linode_token
order by
  age_days desc;
```

### Tokens expiring in the next 30 days
Discover the segments that have tokens expiring in the next 30 days. This is useful for proactively managing and renewing such tokens before they lapse, ensuring uninterrupted access to linked services.

```sql+postgres
select
  token,
  expiry
from
  linode_token
where
  expiry < current_date + interval '30 days';
```

```sql+sqlite
select
  token,
  expiry
from
  linode_token
where
  expiry < date('now','+30 days');
```

### Tokens will full permissions
Explore which tokens have full permissions in Linode. This can be useful for identifying potential security risks and ensuring appropriate access control.

```sql+postgres
select
  *
from
  linode_token
where
  scopes ? '*'
```

```sql+sqlite
Error: SQLite does not support '?' operator for JSON objects.
```

### Scopes by token
Explore which Linode API tokens have been granted specific permissions, as this information can be crucial for managing access rights and ensuring security within your system.

```sql+postgres
select
  t.token,
  scope
from
  linode_token as t,
  jsonb_array_elements_text(t.scopes) as scope;
```

```sql+sqlite
select
  t.token,
  json_extract(scope.value, ')
from
  linode_token as t,
  json_each(t.scopes) as scope;
```