---
title: mconn_snapshot_mount
page_id: schema-mconn-snapshot-mount-fa1dcec9
path: schemas
description: Snapshot Mount
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_mount

Snapshot Mount

```yaml
{"description": "Snapshot Mount", "type": "object", "properties": {"available_bytes": {"description": "Available disk size (bytes)", "type": "number"}, "available_inodes": {"description": "Available inodes on filesystem", "type": "number"}, "file_system": {"description": "File system on disk (EXT4, NTFS, etc.)", "type": "string"}, "is_read_only": {"description": "Determines whether the disk is read-only", "type": "boolean"}, "is_removable": {"description": "Determines whether the disk is removable", "type": "boolean"}, "kind": {"description": "Kind of disk (HDD, SSD, etc.)", "type": "string"}, "mount_point": {"description": "Path where disk is mounted", "type": "string"}, "name": {"description": "Name of the disk mount", "type": "string"}, "total_bytes": {"description": "Total disk size (bytes)", "type": "number"}, "total_inodes": {"description": "Total inodes on filesystem", "type": "number"}}, "required": ["name", "kind", "file_system", "mount_point"]}
```
