---
title: "Steampipe Table: linode_event - Query Linode Events using SQL"
description: "Allows users to query Linode Events, specifically providing insights into system events related to Linode instances."
---

# Table: linode_event - Query Linode Events using SQL

Linode Events are system events that occur on your Linode instances. They can include actions such as booting, shutting down, or resizing a Linode, and can be used to track changes and incidents in your environment. These events provide important information about the activity and performance of your resources.

## Table Usage Guide

The `linode_event` table provides insights into system events related to Linode instances. As a system administrator or DevOps engineer, explore event-specific details through this table, including event types, timestamps, and associated metadata. Utilize it to monitor system activity, track changes, and identify potential issues related to your Linode instances.

## Examples

### List token create events
Explore which events are related to the creation of tokens. This can be useful in understanding and auditing security measures, as token creation often relates to authentication processes.

```sql+postgres
select
  *
from
  linode_event
where
  action = 'token_create';
```

```sql+sqlite
select
  *
from
  linode_event
where
  action = 'token_create';
```

### Events related to a specific domain
Discover the segments that pertain to a specific domain, enabling you to analyze and understand events related to that particular domain.

```sql+postgres
select
  *
from
  linode_event
where
  entity ->> 'type' = 'domain'
  and entity ->> 'label' = 'steampipe.io';
```

```sql+sqlite
select
  *
from
  linode_event
where
  json_extract(entity, '$.type') = 'domain'
  and json_extract(entity, '$.label') = 'steampipe.io';
```