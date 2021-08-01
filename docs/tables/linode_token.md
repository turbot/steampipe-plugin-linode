# Table: linode_token

List tokens for the Linode account.

## Examples

### List tokens

```sql
select
  *
from
  linode_token
```

### Tokens by age in days

```sql
select
  token,
  created,
  date_part('day', age(current_timestamp, created)) as age_days
from
  linode_token
order by
  age desc
```

### Tokens expiring in the next 30 days

```sql
select
  token,
  expiry
from
  linode_token
where
  expiry < current_date + interval '30 days'
```

### Tokens will full permissions

```sql
select
  *
from
  linode_token
where
  scopes ? '*'
```

### Scopes by token

```sql
select
  t.token,
  scope
from
  linode_token as t,
  jsonb_array_elements_text(t.scopes) as scope
```
