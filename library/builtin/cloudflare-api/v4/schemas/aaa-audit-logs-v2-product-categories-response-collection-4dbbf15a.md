---
title: aaa_audit-logs-v2-product-categories-response-collection
page_id: schema-aaa-audit-logs-v2-product-categories-response-collection-4dbbf15a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-product-categories-response-collection

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/aaa_messages-2"}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/aaa_audit-logs-v2-product-category"}}, "success": {"description": "Indicates whether the API call was successful", "type": "boolean", "example": true, "enum": [true]}}}
```
