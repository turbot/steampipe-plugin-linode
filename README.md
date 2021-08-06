![image](https://hub.steampipe.io/images/plugins/turbot/linode-social-graphic.png)

# Linode Plugin for Steampipe

Use SQL to query instances, domains and more from Linode.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/linode)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/linode/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-linode/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install linode
```

Run a query:

```sql
select
  label,
  region,
  status
from
  linode_instance
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-linode.git
cd steampipe-plugin-linode
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/linode.spc
```

Try it!

```
steampipe query
> .inspect linode
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-linode/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Linode Plugin](https://github.com/turbot/steampipe-plugin-linode/labels/help%20wanted)
