---
title: posture-api_RemediationJobSummary
page_id: schema-posture-api-remediationjobsummary-6bf9047c
path: schemas
description: Summary information about a remediation job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobSummary

Summary information about a remediation job.

```yaml
{"description": "Summary information about a remediation job.", "type": "object", "properties": {"created_at": {"description": "When the remediation job was created.", "type": "string", "format": "date-time", "example": "2025-03-18T18:30:15.123456Z"}, "id": {"description": "Unique identifier for the remediation job.", "type": "string", "format": "uuid", "example": "123e4567-e89b-12d3-a456-426614174000"}, "stale": {"description": "Whether this remediation job is stale (created before the finding instance's affliction_date).", "type": "boolean", "example": false}, "status": {"$ref": "#/components/schemas/posture-api_RemediationJobStatusEnum"}}, "example": {"created_at": "2025-03-18T18:30:15.123456Z", "id": "123e4567-e89b-12d3-a456-426614174000", "stale": false, "status": "pending"}, "required": ["id", "status", "created_at", "stale"]}
```
