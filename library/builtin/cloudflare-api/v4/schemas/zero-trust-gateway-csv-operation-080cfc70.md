---
title: zero-trust-gateway_csv_operation
page_id: schema-zero-trust-gateway-csv-operation-080cfc70
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_csv_operation

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "data": {"$ref": "#/components/schemas/zero-trust-gateway_csv_operation_data"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}, "operation_type": {"description": "The type of operation.", "type": "string", "example": "create_list", "enum": ["create_list"]}, "processing_error": {"description": "A human-readable error message if the operation failed. Only present when the operation status is `failed`.", "type": "string", "nullable": true}, "result": {"description": "The result of the operation. Only present when the operation has completed successfully.", "type": "string", "nullable": true}, "status": {"description": "The status of the operation.", "type": "string", "example": "pending", "enum": ["pending", "active", "failed", "complete"]}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}}}
```
