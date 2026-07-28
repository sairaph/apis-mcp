---
title: r2-data-catalog_table-identifier
page_id: schema-r2-data-catalog-table-identifier-ad4bb7bb
path: schemas
description: Specifies a unique table identifier within a catalog.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_table-identifier

Specifies a unique table identifier within a catalog.

```yaml
{"description": "Specifies a unique table identifier within a catalog.", "type": "object", "properties": {"name": {"description": "Specifies the table name.", "type": "string", "example": "events"}, "namespace": {"$ref": "#/components/schemas/r2-data-catalog_namespace-identifier"}}, "required": ["namespace", "name"]}
```
