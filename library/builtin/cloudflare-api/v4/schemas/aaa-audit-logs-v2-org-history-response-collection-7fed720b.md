---
title: aaa_audit-logs-v2-org-history-response-collection
page_id: schema-aaa-audit-logs-v2-org-history-response-collection-7fed720b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-org-history-response-collection

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/aaa_messages-2"}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/aaa_audit-logs-v2-org"}}, "result_info": {"$ref": "#/components/schemas/aaa_audit-logs-v2-history-result-info"}, "success": {"description": "Indicates whether the API call was successful", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "result", "result_info"]}
```
