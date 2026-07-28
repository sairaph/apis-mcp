---
title: posture-api_RemediationJobAsset
page_id: schema-posture-api-remediationjobasset-a4642fa5
path: schemas
description: Asset information for a remediation job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJobAsset

Asset information for a remediation job.

```yaml
{"description": "Asset information for a remediation job.", "type": "object", "properties": {"category": {"$ref": "#/components/schemas/posture-api_RemediationJobAssetCategory"}, "external_id": {"description": "External identifier from the source system.", "type": "string", "example": "c416bc38-75db-425f-ae25-c37b5df5c37f"}, "fields": {"description": "Additional fields associated with the asset.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobAssetField"}}, "id": {"description": "Unique identifier for the asset.", "type": "string", "format": "uuid", "example": "a1b2c3d4-5678-9abc-def0-123456789abc"}, "link": {"description": "Direct link to the asset.", "type": "string", "format": "uri", "example": "https://dashboard.microsoft.com/files/details", "nullable": true}, "name": {"description": "Human-readable name of the asset.", "type": "string", "example": "Microsoft File Publicly Accessible"}}, "required": ["id", "external_id", "name", "fields", "category"]}
```
