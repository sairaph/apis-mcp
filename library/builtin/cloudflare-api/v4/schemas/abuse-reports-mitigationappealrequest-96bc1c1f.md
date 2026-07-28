---
title: abuse-reports_MitigationAppealRequest
page_id: schema-abuse-reports-mitigationappealrequest-96bc1c1f
path: schemas
description: Submit an appeal for a report. Provide either a list of mitigations to appeal, or an appeal type with its supporting details, but not both. When type is "counter_notice", the counter-notice details are required.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationAppealRequest

Submit an appeal for a report. Provide either a list of mitigations to appeal, or an appeal type with its supporting details, but not both. When type is "counter_notice", the counter-notice details are required.

```yaml
{"description": "Submit an appeal for a report. Provide either a list of mitigations to appeal, or an appeal type with its supporting details, but not both. When type is \"counter_notice\", the counter-notice details are required.", "type": "object", "properties": {"appeals": {"description": "List of mitigations to appeal.", "type": "array", "items": {"$ref": "#/components/schemas/abuse-reports_MitigationAppeal"}}, "data": {"$ref": "#/components/schemas/abuse-reports_DMCACounterNotice"}, "type": {"$ref": "#/components/schemas/abuse-reports_AppealType"}}, "oneOf": [{"required": ["appeals"]}, {"required": ["type"]}]}
```
