---
title: aig-billing_GetTopupLimitsResponse
page_id: schema-aig-billing-gettopuplimitsresponse-a43ba0d1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetTopupLimitsResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/aig-billing_GetTopupLimitsResult"}, "result_info": {"type": "object", "properties": {"has_more": {"type": "boolean"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["has_more", "page", "per_page", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "errors", "messages"]}
```
