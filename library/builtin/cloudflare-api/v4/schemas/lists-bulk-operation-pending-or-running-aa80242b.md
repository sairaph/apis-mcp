---
title: lists_bulk_operation_pending_or_running
page_id: schema-lists-bulk-operation-pending-or-running-aa80242b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_bulk_operation_pending_or_running

```yaml
{"type": "object", "properties": {"id": {"$ref": "#/components/schemas/lists_operation_id"}, "status": {"description": "The current status of the asynchronous operation.", "type": "string", "example": "pending", "enum": ["pending", "running"], "readOnly": true, "x-auditable": true}}, "required": ["id", "status"]}
```
