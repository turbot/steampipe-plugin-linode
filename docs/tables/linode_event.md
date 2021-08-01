# Table: linode_event

Query events from your Linode account.

## Examples

### List token create events

```sql
select
  *
from
  linode_event
where
  action = 'token_create'
```

### Events related to a specific domain

```sql
select
  *
from
  linode_event
where
  entity ->> 'type' = 'domain'
  and entity ->> 'label' = 'steampipe.io'
```
