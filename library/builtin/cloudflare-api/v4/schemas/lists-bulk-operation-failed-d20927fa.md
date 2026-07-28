---
title: lists_bulk_operation_failed
page_id: schema-lists-bulk-operation-failed-d20927fa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_bulk_operation_failed

```yaml
{"type": "object", "properties": {"completed": {"$ref": "#/components/schemas/lists_completed"}, "error": {"description": "A message describing the error when the status is `failed`.", "type": "string", "example": "This list is at the maximum number of items", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/lists_operation_id"}, "status": {"description": "The current status of the asynchronous operation.", "type": "string", "example": "failed", "enum": ["failed"], "readOnly": true, "x-auditable": true}}, "required": ["id", "status", "error", "completed"]}
```
