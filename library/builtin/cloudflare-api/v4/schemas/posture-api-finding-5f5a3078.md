---
title: posture-api_Finding
page_id: schema-posture-api-finding-5f5a3078
path: schemas
description: Aggregated finding information with counts and metadata. This is optimized for list API queries and represents a finding along with its instance statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_Finding

Aggregated finding information with counts and metadata. This is optimized for list API queries and represents a finding along with its instance statistics.

```yaml
{"description": "Aggregated finding information with counts and metadata. This is optimized for list API queries and represents a finding along with its instance statistics.", "type": "object", "properties": {"active_count": {"description": "Number of active problematic instances identified in the security finding.", "type": "integer", "example": 5, "readOnly": true}, "archived_count": {"description": "Number of archived instances identified in the security finding.", "type": "integer", "example": 2, "readOnly": true}, "finding": {"$ref": "#/components/schemas/posture-api_FindingType"}, "id": {"description": "Base64 encoded identifier of the security finding.", "type": "string", "format": "byte", "example": "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAxOjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMgo=", "readOnly": true}, "ignored": {"description": "Determines if finding is currently ignored.", "type": "boolean", "example": false, "readOnly": true}, "instance_count": {"description": "Number of total (Active or archived) problematic instances identified in the security finding.", "type": "integer", "example": 7, "readOnly": true}, "integration": {"$ref": "#/components/schemas/posture-api_IntegrationSummary"}, "latest_affliction_date": {"description": "Timestamp of the latest affliction date of an active finding.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.700131Z", "readOnly": true}, "severity_override": {"$ref": "#/components/schemas/posture-api_FindingSeverityOverride"}}, "required": ["active_count", "archived_count", "finding", "id", "ignored", "instance_count", "integration", "latest_affliction_date"]}
```
