---
title: r2-data-catalog_snapshot-expiration-update-params
page_id: schema-r2-data-catalog-snapshot-expiration-update-params-90cb960a
path: schemas
description: Updates snapshot expiration configuration (all fields optional).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_snapshot-expiration-update-params

Updates snapshot expiration configuration (all fields optional).

```yaml
{"description": "Updates snapshot expiration configuration (all fields optional).", "type": "object", "properties": {"max_snapshot_age": {"description": "Updates the maximum age for snapshots optionally.", "type": "string", "pattern": "^\\d+(d|h|m|s)$"}, "min_snapshots_to_keep": {"description": "Updates the minimum number of snapshots to retain optionally.", "type": "integer", "format": "int64", "minimum": 1}, "state": {"description": "Updates the state optionally.", "allOf": [{"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-state"}]}}}
```
