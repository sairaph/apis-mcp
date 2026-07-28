---
title: posture-api_ContentAsset
page_id: schema-posture-api-contentasset-fa1c65c3
path: schemas
description: Content asset with DLP information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_ContentAsset

Content asset with DLP information.

```yaml
{"description": "Content asset with DLP information.", "type": "object", "properties": {"asset_id": {"description": "Unique identifier for the asset.", "type": "string", "format": "uuid", "example": "e6910838-4b91-45e9-b2b4-91bb23cb9762"}, "asset_name": {"description": "Name of the asset.", "type": "string", "example": "Test Asset Name"}, "dlp_contexts": {"description": "DLP context information for this asset.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_DlpContext"}}, "dlp_profile_count": {"description": "Number of DLP profiles that flagged this asset.", "type": "integer", "example": 2}, "dlp_profile_ids": {"description": "IDs of DLP profiles that flagged this asset.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["c12f2059-8df4-43f8-9eb9-d27112d92b63", "822c051b-0bb4-4747-8929-471a1d506eef"]}, "integration": {"$ref": "#/components/schemas/posture-api_IntegrationSummary"}, "latest_affliction_date": {"description": "Most recent date this asset was flagged.", "type": "string", "format": "date-time", "example": "2024-10-18T19:53:57.626659Z"}}, "required": ["asset_id", "asset_name", "latest_affliction_date", "dlp_profile_ids", "dlp_profile_count", "dlp_contexts", "integration"]}
```
