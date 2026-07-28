---
title: posture-api_RemediationJobFailure
page_id: schema-posture-api-remediationjobfailure-98f2bf18
path: schemas
description: Information about a failed remediation job creation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobFailure

Information about a failed remediation job creation.

```yaml
{"description": "Information about a failed remediation job creation.", "type": "object", "properties": {"error": {"description": "Error message describing the failure.", "type": "string", "example": "Failed to create remediation job"}, "finding_instance_id": {"description": "ID of the finding instance that failed to create a remediation job.", "type": "string", "format": "uuid", "example": "2e6b4c8a-9d1f-4e3b-8c7a-5f9e2d1a6b4c"}}, "required": ["finding_instance_id", "error"]}
```
