---
title: posture-api_RemediationType
page_id: schema-posture-api-remediationtype-e1e229bf
path: schemas
description: Information about a remediation type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationType

Information about a remediation type.

```yaml
{"description": "Information about a remediation type.", "type": "object", "properties": {"description": {"description": "A description of the action(s) taken by the remediation type.", "type": "string", "example": "Remove publicly accessible URL granting edit access"}, "display_name": {"description": "The name of the remediation type as displayed in the cloudflare dashboard.", "type": "string", "example": "Remove Publicly Accessible URL - Edit Access"}, "finding_type_id": {"description": "The identifier of the finding_type which this remediation type should remediate.", "type": "string", "format": "uuid", "example": "6a790513-bbb5-4933-8971-76a744ec5448"}, "id": {"description": "The identifier for the remediation type.", "type": "string", "format": "uuid", "example": "7d736ac5-ed3b-46d5-9375-7025175ba1d9"}, "remediation_type": {"description": "The name of the remediation type.", "type": "string", "example": "Microsoft: Remove Publicly Accessible URL - Edit Access"}}, "required": ["id", "finding_type_id", "remediation_type", "display_name", "description"]}
```
