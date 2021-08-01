# Table: linode_region

List regions for the Linode account.

## Examples

### List all regions

```sql
select
  *
from
  linode_region
order by
  id
```

### List all US regions

```sql
select
  *
from
  linode_region
where
  country = 'us'
order by
  id
```
