# Table: linode_instance

List Linodes (instances) across all regions in your account.

## Examples

### List all instances

```sql
select
  *
from
  linode_instance
```

### Instances by region

```sql
select
  region,
  count(*)
from
  linode_instance
group by
  region
```

### Instances by type

```sql
select
  instance_type,
  count(*)
from
  linode_instance
group by
  instance_type
```

### Instances with a given tag

```sql
select
  label,
  status,
  tags
from
  linode_instance
where
  tags ? 'foo'
```

### Running instances

```sql
select
  label,
  ipv4,
  status,
  tags
from
  linode_instance
where
  status = 'running'
```
