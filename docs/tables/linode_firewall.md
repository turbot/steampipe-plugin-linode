---
title: "Steampipe Table: linode_firewall - Query Linode Firewalls using SQL"
description: "Allows users to query Linode Firewalls, providing detailed information about each firewall's configuration, status, and rules."
---

# Table: linode_firewall - Query Linode Firewalls using SQL

Linode Firewalls provide a layer of security by controlling the traffic to and from your Linode instances. They allow you to define rules that either allow or block traffic based on specific criteria, helping you protect your infrastructure from unauthorized access and attacks.

## Table Usage Guide

The `linode_firewall` table allows system administrators and DevOps engineers to query and manage the firewalls associated with their Linode instances. Through this table, you can explore firewall-specific details such as creation and update times, current status, and associated rules.

## Examples

### Basic info
Retrieve a list of all firewalls in your Linode account to get an overview of your security infrastructure.

```sql+postgres
select
  id,
  created,
  updated,
  label
from
  linode_firewall;
```

```sql+sqlite
select
  id,
  created,
  updated,
  label
from
  linode_firewall;
```

### Firewalls that are enabled
Explore the enabled firewalls to ensure that your security configurations are active and properly managed.

```sql+postgres
select
  label,
  status,
  created,
  updated
from
  linode_firewall
where
  status = 'enabled';
```

```sql+sqlite
select
  label,
  status,
  created,
  updated
from
  linode_firewall
where
  status = 'enabled';
```

### Recently updated firewalls
List all firewalls that have been updated in the last 30 days to track recent changes in your security settings.

```sql+postgres
select
  label,
  updated,
  status
from
  linode_firewall
where
  updated >= current_date - interval '30 days';
```

```sql+sqlite
select
  label,
  updated,
  status
from
  linode_firewall
where
  updated >= date('now', '-30 days');
```

### Get inbound rule details of firewalls
This query retrieves detailed information about the inbound rules configured for each Linode firewall, including label, ports, action, protocol, description, and addresses.

```sql+postgres
select
  f.id,
  f.label,
  i ->> 'label' as inbound_label,
  i ->> 'ports' as inbound_ports,
  i ->> 'action' as inbound_action,
  i ->> 'protocol' as inbound_protocol,
  i ->> 'description' as inbound_description,
  i -> 'addresses' as inbound_addresses
from
  linode_firewall as f,
  jsonb_array_elements(rules -> 'inbound') as i;
```

```sql+sqlite
select
  f.id,
  f.label,
  json_extract(i.value, '$.label') as inbound_label,
  json_extract(i.value, '$.ports') as inbound_ports,
  json_extract(i.value, '$.action') as inbound_action,
  json_extract(i.value, '$.protocol') as inbound_protocol,
  json_extract(i.value, '$.description') as inbound_description,
  json_extract(i.value, '$.addresses') as inbound_addresses
from
  linode_firewall as f,
  json_each(rules -> 'inbound') as i;
```

### Get outbound rule details of firewalls
This query retrieves detailed information about the outbound rules configured for each Linode firewall, including label, ports, action, protocol, description, and addresses.

```sql+postgres
select
  f.id,
  f.label,
  o ->> 'label' as outbound_label,
  o ->> 'ports' as outbound_ports,
  o ->> 'action' as outbound_action,
  o ->> 'protocol' as outbound_protocol,
  o ->> 'description' as outbound_description,
  o -> 'addresses' as outbound_addresses
from
  linode_firewall as f,
  jsonb_array_elements(rules -> 'outbound') as o;
```

```sql+sqlite
select
  f.id,
  f.label,
  json_extract(o.value, '$.label') as outbound_label,
  json_extract(o.value, '$.ports') as outbound_ports,
  json_extract(o.value, '$.action') as outbound_action,
  json_extract(o.value, '$.protocol') as outbound_protocol,
  json_extract(o.value, '$.description') as outbound_description,
  json_extract(o.value, '$.addresses') as outbound_addresses
from
  linode_firewall as f,
  json_each(rules -> 'outbound') as o;
```