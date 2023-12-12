---
title: "Steampipe Table: linode_profile - Query Linode Profiles using SQL"
description: "Allows users to query Linode Profiles, specifically user's profile information and settings, providing insights into account management and user preferences."
---

# Table: linode_profile - Query Linode Profiles using SQL

Linode Profiles are a part of the Linode API that allows users to manage their account settings and preferences. The profile encompasses a user's contact information, two-factor authentication settings, authorized third-party applications, and more. It is a centralized way to manage user-specific details within the Linode environment.

## Table Usage Guide

The `linode_profile` table provides insights into user profiles within Linode. As a system administrator, explore user-specific details through this table, including contact information, two-factor authentication settings, and authorized applications. Utilize it to manage user preferences, such as timezone settings, authentication methods, and to monitor authorized third-party applications.

## Examples

### Get profile data
Explore your Linode profile's overall configuration and settings to gain insights into your current setup and identify areas for potential improvement.

```sql+postgres
select
  *
from
  linode_profile;
```

```sql+sqlite
select
  *
from
  linode_profile;
```