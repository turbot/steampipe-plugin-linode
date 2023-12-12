---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/linode.svg"
brand_color: "#00b050"
display_name: "Linode"
short_name: "linode"
description: "Steampipe plugin to query resources, users and more from Linode."
og_description: "Query Linode with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/linode-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Linode + Steampipe

[Linode](https://linode.com) is a cloud hosting company that provides virtual private servers and other infrastructure services.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List instances in your Linode account:

```sql
select
  label,
  region,
  status
from
  linode_instance
```

```
+-------------+---------+---------+
| label       | region  | status  |
+-------------+---------+---------+
| my-instance | us-east | running |
+-------------+---------+---------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/linode/tables)**

## Get started

### Install

Download and install the latest Linode plugin:

```bash
steampipe plugin install linode
```

### Credentials

No credentials are required.

### Configuration

Installing the latest linode plugin will create a config file (`~/.steampipe/config/linode.spc`) with a single connection named `linode`:

```hcl
connection "linode" {
  plugin = "linode"
  token  = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fbd"
}
```

- `token` - API token from Linode.

## Multi-Account Connections

You may create multiple linode connections:

```hcl
connection "linode_dev" {
  plugin    = "linode"
  token     = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9965d91fbd"
}

connection "linode_qa" {
  plugin    = "linode"
  token     = "5a76843869c183a4ty001c79102bfa1f667f39a2ea0ba857c9a35a9965d91fbd"
}

connection "linode_prod" {
  plugin    = "linode"
  token     = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a7765d91fbd"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from linode_qa.linode_user
```

You can create multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "linode_all" {
  plugin      = "linode"
  type        = "aggregator"
  connections = ["linode_dev", "linode_qa", "linode_prod"]
}
```

Querying tables from this connection will return results from the `linode_dev`, `linode_qa`, and `linode_prod` connections:

```sql
select * from linode_all.linode_user
```

Alternatively, you can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically so that it is the first connection in the search path (i.e. `linode_all` comes before `linode_dev`):

```sql
select * from linode_user
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the linode_user plugin connections whose names begin with `linode_`:

```hcl
connection "linode_all" {
  type        = "aggregator"
  plugin      = "linode_user"
  connections = ["linode_*"]
}
```


