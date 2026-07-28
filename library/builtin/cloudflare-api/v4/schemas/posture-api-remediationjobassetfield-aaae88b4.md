---
title: posture-api_RemediationJobAssetField
page_id: schema-posture-api-remediationjobassetfield-aaae88b4
path: schemas
description: Additional field information for a remediation job asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobAssetField

Additional field information for a remediation job asset.

```yaml
{"description": "Additional field information for a remediation job asset.", "type": "object", "properties": {"link": {"description": "Optional link associated with the field.", "type": "string", "format": "uri", "example": "https://dashboard.microsoft.com/files/details", "nullable": true}, "name": {"description": "Field name.", "type": "string", "example": "File Name"}, "value": {"description": "Field value (can be string, number, or boolean).", "type": "string", "example": "sensitive-document.xlsx", "oneOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "required": ["name", "value"]}
```
