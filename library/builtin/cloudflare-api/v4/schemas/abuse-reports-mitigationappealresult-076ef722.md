---
title: abuse-reports_MitigationAppealResult
page_id: schema-abuse-reports-mitigationappealresult-076ef722
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationAppealResult

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_MitigationListItem"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}, "total_pages": {"type": "number"}}, "required": ["page", "per_page", "count", "total_count", "total_pages"]}, "success": {"type": "boolean"}}, "required": ["success", "result_info"]}
```
