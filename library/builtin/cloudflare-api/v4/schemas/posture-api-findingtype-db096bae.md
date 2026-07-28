---
title: posture-api_FindingType
page_id: schema-posture-api-findingtype-db096bae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingType

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_BaseFindingType"}, {"properties": {"description": {"description": "Detailed description of the finding.", "type": "string", "example": "This finding indicates that a file in your Slack workspace is publicly accessible.", "nullable": true}, "remediation": {"$ref": "#/components/schemas/posture-api_FindingRemediation"}}, "type": "object"}], "required": ["category", "id", "name", "remediation", "severity", "vendor"]}
```
