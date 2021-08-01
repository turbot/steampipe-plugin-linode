# Table: linode_account

Get account information, including balances.

## Examples

### Get account information

```sql
select
  *
from
  linode_account
```

### Balances

```sql
select
  email,
  balance,
  balance_uninvoiced,
  credit_card
from
  linode_account
```
