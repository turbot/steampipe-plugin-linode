# Table: linode_type

List instance types for the Linode account.

## Examples

### List all types

```sql
select
  *
from
  linode_type
order by
  price_monthly
```

### List all dedicated instance types

```sql
select
  *
from
  linode_type
where
  class = 'dedicated'
order by
  price_monthly
```
