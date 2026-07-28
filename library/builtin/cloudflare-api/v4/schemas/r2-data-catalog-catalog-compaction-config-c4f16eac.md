---
title: r2-data-catalog_catalog-compaction-config
page_id: schema-r2-data-catalog-catalog-compaction-config-c4f16eac
path: schemas
description: Configures compaction for catalog maintenance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog-compaction-config

Configures compaction for catalog maintenance.

```yaml
{"description": "Configures compaction for catalog maintenance.", "type": "object", "properties": {"state": {"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-state"}, "target_size_mb": {"$ref": "#/components/schemas/r2-data-catalog_catalog-target-file-size"}}, "required": ["state", "target_size_mb"]}
```
