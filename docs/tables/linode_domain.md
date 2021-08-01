# Table: linode_domain

List domains for the Linode account.

## Examples

### List all domains

```sql
select
  *
from
  linode_domain
```

### Domains with a given tag

```sql
select
  domain,
  tags
from
  linode_volume
where
  tags ? 'foo'
```
