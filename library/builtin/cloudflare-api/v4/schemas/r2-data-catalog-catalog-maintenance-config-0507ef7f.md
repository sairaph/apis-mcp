---
title: r2-data-catalog_catalog-maintenance-config
page_id: schema-r2-data-catalog-catalog-maintenance-config-0507ef7f
path: schemas
description: Configures maintenance for the catalog.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog-maintenance-config

Configures maintenance for the catalog.

```yaml
{"description": "Configures maintenance for the catalog.", "type": "object", "properties": {"compaction": {"$ref": "#/components/schemas/r2-data-catalog_catalog-compaction-config"}, "snapshot_expiration": {"$ref": "#/components/schemas/r2-data-catalog_snapshot-expiration-config"}}}
```
