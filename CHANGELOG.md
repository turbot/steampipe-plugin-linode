## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#65](https://github.com/turbot/steampipe-plugin-linode/pull/65))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#65](https://github.com/turbot/steampipe-plugin-linode/pull/65))

## v0.7.0 [2024-08-08]

_What's new?_

- New tables added
  - [linode_firewall](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_firewall)
  - [linode_node_balancer](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_node_balancer)

_Enhancements_

- The `euuid` column has now been assigned as a connection key column across all the tables which facilitates more precise and efficient querying across multiple Linode accounts. ([#56](https://github.com/turbot/steampipe-plugin-linode/pull/56))
- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package. ([#60](https://github.com/turbot/steampipe-plugin-linode/pull/60))
- Added the `version` flag to the plugin's Export tool. ([#65](https://github.com/turbot/steampipe-export/pull/65))

_Dependencies_

- Recompiled plugin with [linode-sdk-for-go v1.37.0](https://github.com/linode/linodego/releases/tag/v1.37.0). ([#56](https://github.com/turbot/steampipe-plugin-linode/pull/56))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5101-2024-05-09) which ensures that `QueryData` passed to `ConnectionKeyColumns` value callback is populated with `ConnectionManager`. ([#55](https://github.com/turbot/steampipe-plugin-linode/pull/55))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#47](https://github.com/turbot/steampipe-plugin-linode/pull/47))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#47](https://github.com/turbot/steampipe-plugin-linode/pull/47))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-linode/blob/main/docs/LICENSE). ([#47](https://github.com/turbot/steampipe-plugin-linode/pull/47))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#46](https://github.com/turbot/steampipe-plugin-linode/pull/46))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#31](https://github.com/turbot/steampipe-plugin-linode/pull/31))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#28](https://github.com/turbot/steampipe-plugin-linode/pull/28))
- Recompiled plugin with Go version `1.21`. ([#28](https://github.com/turbot/steampipe-plugin-linode/pull/28))

## v0.4.0 [2023-07-17]

_Enhancements_

- Updated the `docs/index.md` file to include multi-account configuration examples. ([#19](https://github.com/turbot/steampipe-plugin-linode/pull/19))

## v0.3.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#17](https://github.com/turbot/steampipe-plugin-linode/pull/17))

## v0.2.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#14](https://github.com/turbot/steampipe-plugin-linode/pull/14))
- Recompiled plugin with Go version `1.19`. ([#14](https://github.com/turbot/steampipe-plugin-linode/pull/14))

## v0.1.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#10](https://github.com/turbot/steampipe-plugin-linode/pull/10))

## v0.1.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#8](https://github.com/turbot/steampipe-plugin-linode/pull/8))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#7](https://github.com/turbot/steampipe-plugin-linode/pull/7))

## v0.0.2 [2021-11-24]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#4](https://github.com/turbot/steampipe-plugin-linode/pull/4))

## v0.0.1 [2021-08-05]

_What's new?_

- New tables added
  - [linode_account](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_account)
  - [linode_bucket](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_bucket)
  - [linode_domain](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_domain)
  - [linode_domain_record](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_domain_record)
  - [linode_event](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_event)
  - [linode_image](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_image)
  - [linode_instance](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_instance)
  - [linode_kubernetes_cluster](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_kubernetes_cluster)
  - [linode_profile](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_profile)
  - [linode_region](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_region)
  - [linode_tag](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_tag)
  - [linode_token](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_token)
  - [linode_type](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_type)
  - [linode_user](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_user)
  - [linode_volume](https://hub.steampipe.io/plugins/turbot/linode/tables/linode_volume)
