# Table: linode_domain_record

List domain records.

Note: A `domain_id` must be provided in all queries to this table.

## Examples

### List all records for all domains

```sql
select
  d.domain,
  dr.record_type,
  dr.name,
  dr.target
from
  linode_domain as d,
  linode_domain_record as dr
where
  dr.domain_id = d.id
```

### List all domain records for a domain

```sql
select
  *
from
  linode_domain_record
where
  domain_id = 1234
```
