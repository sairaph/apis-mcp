---
title: lists_bulk_operation_completed
page_id: schema-lists-bulk-operation-completed-a9bf9f75
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_bulk_operation_completed

```yaml
{"type": "object", "properties": {"completed": {"$ref": "#/components/schemas/lists_completed"}, "id": {"$ref": "#/components/schemas/lists_operation_id"}, "status": {"description": "The current status of the asynchronous operation.", "type": "string", "example": "completed", "enum": ["completed"], "readOnly": true, "x-auditable": true}}, "required": ["id", "status", "completed"]}
```
