---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/linode.svg"
brand_color: "#4e7e14"
display_name: "Linode"
short_name: "linode"
description: "Steampipe plugin to query resources, users and more from Linode."
og_description: "Query Linode with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/linode-social-graphic.png"
---

# Linode + Steampipe

[Linode](https://linode.com) is a cloud hosting company that provides virtual private servers and other infrastructure services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-linode
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
