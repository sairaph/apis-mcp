---
title: posture-api_RemediationJobsExportOrderFilterRequest
page_id: schema-posture-api-remediationjobsexportorderfilterrequest-65ff9016
path: schemas
description: Order specification for remediation jobs exports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobsExportOrderFilterRequest

Order specification for remediation jobs exports.

```yaml
{"description": "Order specification for remediation jobs exports.", "type": "object", "properties": {"direction": {"$ref": "#/components/schemas/posture-api_DirectionEnum"}, "name": {"description": "Which field to use when ordering the remediation jobs.", "type": "string", "example": "last_updated_at", "enum": ["asset_name", "finding_type_name", "integration_name", "status", "last_updated_at", "affliction_date"]}}, "required": ["direction", "name"]}
```
