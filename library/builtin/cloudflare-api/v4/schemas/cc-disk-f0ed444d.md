---
title: cc_Disk
page_id: schema-cc-disk-f0ed444d
path: schemas
description: The disk configuration for this deployment. By default, all containers have a disk size of 2GB.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_Disk

The disk configuration for this deployment. By default, all containers have a disk size of 2GB.

```yaml
{"description": "The disk configuration for this deployment. By default, all containers have a disk size of 2GB.", "type": "object", "properties": {"size": {"$ref": "#/components/schemas/cc_DiskSizeWithUnit"}, "size_mb": {"description": "Size of the disk, in MB.", "type": "integer"}}}
```
