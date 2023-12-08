---
title: "Steampipe Table: linode_image - Query Linode Images using SQL"
description: "Allows users to query Linode Images, providing detailed information about the available custom images on a user's Linode account."
---

# Table: linode_image - Query Linode Images using SQL

Linode Images allows users to capture and deploy disk images. These images can be used to create Linode instances or to resize existing Linodes. Linode Images are a great way to back up your data or to create templates for regularly-used configurations.

## Table Usage Guide

The `linode_image` table provides insights into custom images within Linode. As a DevOps engineer, explore image-specific details through this table, including labels, descriptions, and associated metadata. Utilize it to uncover information about images, such as their creation date, size, type, and whether they are public or private.

## Examples

### All images
Explore all available system images on your Linode account. This is useful to manage and keep track of the different operating systems and configurations you have available for deployment.

```sql+postgres
select
  *
from
  linode_image;
```

```sql+sqlite
select
  *
from
  linode_image;
```

### Ubuntu images
Explore which Linode images are related to Ubuntu. This query is useful for identifying all Ubuntu-related images in your Linode environment, helping you manage your resources more effectively.

```sql+postgres
select
  *
from
  linode_image
where
  id like '%ubuntu%';
```

```sql+sqlite
select
  *
from
  linode_image
where
  id like '%ubuntu%';
```

### Image information for each instances
Gain insights into the specific image details associated with each instance, enabling you to understand the underlying configurations. This is useful for ensuring that instances are running on the correct images, thereby improving system consistency and reliability.

```sql+postgres
select
  i.label as instance_label,
  im.*
from
  linode_instance as i,
  linode_image as im
where
  i.image = im.id;
```

```sql+sqlite
select
  i.label as instance_label,
  im.*
from
  linode_instance as i
join
  linode_image as im
on
  i.image = im.id;
```