---
title: posture-api_RemediationJobsExportFilterRequest
page_id: schema-posture-api-remediationjobsexportfilterrequest-0b43706c
path: schemas
description: Filter specification for remediation jobs export.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobsExportFilterRequest

Filter specification for remediation jobs export.

```yaml
{"description": "Filter specification for remediation jobs export.", "type": "object", "properties": {"integration_id": {"description": "Filter by multiple integration IDs.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["55d7337e-1d0a-49fc-9826-925ba40df035"]}, "max_updated_at": {"description": "Filter to view remediation jobs updated on or before this datetime. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2025-01-01T00:00:00Z"}, "min_updated_at": {"description": "Filter to view remediation jobs updated on or after this datetime. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}, "orders": {"description": "Ordering specifications for the export.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobsExportOrderFilterRequest"}, "default": []}, "search": {"description": "A search term.", "type": "string", "example": "public access"}, "status": {"description": "Filter by remediation job status.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobStatusEnum"}, "example": ["pending", "completed"]}}}
```
