---
title: r2-data-catalog_snapshot-expiration-config
page_id: schema-r2-data-catalog-snapshot-expiration-config-bfcfef23
path: schemas
description: Configures snapshot expiration settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_snapshot-expiration-config

Configures snapshot expiration settings.

```yaml
{"description": "Configures snapshot expiration settings.", "type": "object", "properties": {"max_snapshot_age": {"description": "Specifies the maximum age for snapshots. The system deletes snapshots older than this age.\nFormat: <number><unit> where unit is d (days), h (hours), m (minutes), or s (seconds).\nExamples: \"7d\" (7 days), \"48h\" (48 hours), \"2880m\" (2,880 minutes).\nDefaults to \"7d\".\n", "type": "string", "example": "7d", "pattern": "^\\d+(d|h|m|s)$"}, "min_snapshots_to_keep": {"description": "Specifies the minimum number of snapshots to retain. Defaults to 100.", "type": "integer", "format": "int64", "example": 100, "minimum": 1}, "state": {"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-state"}}, "required": ["state", "min_snapshots_to_keep", "max_snapshot_age"]}
```
