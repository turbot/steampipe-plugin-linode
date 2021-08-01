# Table: linode_bucket

List object storage buckets in the Linode account.

## Examples

### List buckets

```sql
select
  *
from
  linode_bucket
```

### Buckets in us-east-1

```sql
select
  *
from
  linode_bucket
where
  cluster = 'us-east-1'
```
