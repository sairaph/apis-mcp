---
title: abuse-reports_AppealEligibility
page_id: schema-abuse-reports-appealeligibility-85fbeada
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_AppealEligibility

```yaml
{"type": "object", "properties": {"appeal_count": {"description": "Number of appeals submitted against the report so far.", "type": "integer"}, "appealable": {"description": "Whether the report can currently be appealed.", "type": "boolean"}, "has_appealable_mitigations": {"description": "Whether the report has at least one mitigation an appeal could reverse.", "type": "boolean"}, "has_open_appeal": {"description": "Whether the report already has an open (undecided) appeal.", "type": "boolean"}, "max_appeals": {"description": "Maximum number of appeals allowed per report.", "type": "integer"}}, "required": ["appealable", "has_appealable_mitigations", "has_open_appeal", "appeal_count", "max_appeals"]}
```
