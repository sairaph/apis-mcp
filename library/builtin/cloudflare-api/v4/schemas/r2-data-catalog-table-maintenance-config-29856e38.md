---
title: r2-data-catalog_table-maintenance-config
page_id: schema-r2-data-catalog-table-maintenance-config-29856e38
path: schemas
description: Configures maintenance for the table.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_table-maintenance-config

Configures maintenance for the table.

```yaml
{"description": "Configures maintenance for the table.", "type": "object", "properties": {"compaction": {"$ref": "#/components/schemas/r2-data-catalog_table-compaction-config"}, "snapshot_expiration": {"$ref": "#/components/schemas/r2-data-catalog_snapshot-expiration-config"}}}
```
