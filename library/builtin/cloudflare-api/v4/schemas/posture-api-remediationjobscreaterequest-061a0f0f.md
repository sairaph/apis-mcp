---
title: posture-api_RemediationJobsCreateRequest
page_id: schema-posture-api-remediationjobscreaterequest-061a0f0f
path: schemas
description: Request body for creating remediation jobs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobsCreateRequest

Request body for creating remediation jobs.

```yaml
{"description": "Request body for creating remediation jobs.", "type": "object", "properties": {"finding_instance_ids": {"description": "UUIDs identifying Finding Instances.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e", "ab7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d7g"]}, "remediation_type_id": {"description": "A UUID identifying this Remediation Type.", "type": "string", "format": "uuid", "example": "5a7d9e2f-1b3c-4d5e-8f6a-7b8c9d0e1f2a"}}, "required": ["finding_instance_ids", "remediation_type_id"]}
```
