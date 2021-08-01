# Table: linode_image

List images for the Linode account.

## Examples

### All images

```sql
select
  *
from
  linode_image
```

### Ubuntu images

```sql
select
  *
from
  linode_image
where
  id like '%ubuntu%'
```

### Image information for each instances

```sql
select
  i.label as instance_label,
  im.*
from
  linode_instance as i,
  linode_image as im
where
  i.image = im.id
```
