---
title: "Steampipe Table: linode_account - Query Linode Accounts using SQL"
description: "Allows users to query Linode Accounts, providing details about the account, including balance, transfer, and billing information."
---

# Table: linode_account - Query Linode Accounts using SQL

Linode is a cloud hosting provider that offers high-performance SSD Linux servers for all of your infrastructure needs. Linode has received multiple awards for being one of the best virtual private server providers. Linode's Account service provides detailed information about the account, including balance, transfer, and billing information.

## Table Usage Guide

The `linode_account` table provides insights into the account details within Linode. As a cloud engineer, explore account-specific details through this table, including balance, transfer, and billing information. Utilize it to uncover information about the account, such as current balance, pending charges, and the account's last payment date.

## Examples

### Get account information
Review the configuration for your Linode account to gain insights into its settings and details. This is useful for understanding the state and configuration of your account as a whole.

```sql+postgres
select
  *
from
  linode_account;
```

```sql+sqlite
select
  *
from
  linode_account;
```

### Balances
Explore the financial status of your Linode accounts. This query provides an overview of the account balances, uninvoiced balances, and associated credit card details, helping you manage your resources effectively.

```sql+postgres
select
  email,
  balance,
  balance_uninvoiced,
  credit_card
from
  linode_account;
```

```sql+sqlite
select
  email,
  balance,
  balance_uninvoiced,
  credit_card
from
  linode_account;
```