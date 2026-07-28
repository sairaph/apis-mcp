---
title: abuse-reports_MitigationAppeal
page_id: schema-abuse-reports-mitigationappeal-11dfdb5c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationAppeal

```yaml
{"type": "object", "properties": {"id": {"description": "ID of the mitigation to appeal.", "type": "string"}, "reason": {"$ref": "#/components/schemas/abuse-reports_AppealReason"}}, "required": ["id", "reason"]}
```
