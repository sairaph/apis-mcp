---
title: r2-data-catalog_namespace-details
page_id: schema-r2-data-catalog-namespace-details-9c4f5036
path: schemas
description: Contains namespace with metadata details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_namespace-details

Contains namespace with metadata details.

```yaml
{"description": "Contains namespace with metadata details.", "type": "object", "properties": {"created_at": {"description": "Indicates the creation timestamp in ISO 8601 format.", "type": "string", "format": "date-time", "nullable": true}, "namespace": {"$ref": "#/components/schemas/r2-data-catalog_namespace-identifier"}, "namespace_uuid": {"description": "Contains the UUID that persists across renames.", "type": "string", "format": "uuid"}, "updated_at": {"description": "Shows the last update timestamp in ISO 8601 format. Null if never updated.", "type": "string", "format": "date-time", "nullable": true}}, "required": ["namespace", "namespace_uuid"]}
```
