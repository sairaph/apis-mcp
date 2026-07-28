---
title: r2-data-catalog_compaction-update-params
page_id: schema-r2-data-catalog-compaction-update-params-9996edd8
path: schemas
description: Updates compaction configuration (all fields optional).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_compaction-update-params

Updates compaction configuration (all fields optional).

```yaml
{"description": "Updates compaction configuration (all fields optional).", "type": "object", "properties": {"state": {"description": "Updates the state optionally.", "allOf": [{"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-state"}]}, "target_size_mb": {"description": "Updates the target file size optionally.", "allOf": [{"$ref": "#/components/schemas/r2-data-catalog_catalog-target-file-size"}]}}}
```
