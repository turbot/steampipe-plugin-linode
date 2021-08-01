# Table: linode_volume

List volumes for the Linode account.

## Examples

### List volumes

```sql
select
  *
from
  linode_volume
```

### Find volumes in a bad state

```sql
select
  label,
  size,
  status,
  region
from
  linode_volume
where
  status = 'contact_support'
```

### Volumes with a given tag

```sql
select
  label,
  size,
  status,
  region
from
  linode_volume
where
  tags ? 'foo'
```

### Top 5 volumes by size

```sql
select
  label,
  size,
  status,
  region
from
  linode_volume
order by
  size desc
limit
  5
```
