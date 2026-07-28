---
title: aaa_audit_logs_response_collection
page_id: schema-aaa-audit-logs-response-collection-a9fdf731
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit_logs_response_collection

```yaml
{"oneOf": [{"properties": {"errors": {"$ref": "#/components/schemas/aaa_messages"}, "messages": {"$ref": "#/components/schemas/aaa_messages"}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/aaa_audit-logs"}}, "success": {"type": "boolean", "example": true}}}, {"$ref": "#/components/schemas/aaa_api-response-common"}]}
```
