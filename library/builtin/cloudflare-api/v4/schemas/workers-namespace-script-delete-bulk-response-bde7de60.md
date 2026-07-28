---
title: workers_namespace-script-delete-bulk-response
page_id: schema-workers-namespace-script-delete-bulk-response-bde7de60
path: schemas
description: Detail about bulk deletion of scripts in a namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_namespace-script-delete-bulk-response

Detail about bulk deletion of scripts in a namespace.

```yaml
{"description": "Detail about bulk deletion of scripts in a namespace.", "type": "object", "properties": {"deleted": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/workers_uuid"}}, "type": "object"}}, "deleted_count": {"type": "integer", "example": 100}, "has_more": {"type": "boolean", "example": false}}}
```
