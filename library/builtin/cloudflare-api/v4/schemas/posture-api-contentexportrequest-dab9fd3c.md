---
title: posture-api_ContentExportRequest
page_id: schema-posture-api-contentexportrequest-dab9fd3c
path: schemas
description: Request body for creating content exports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_ContentExportRequest

Request body for creating content exports.

```yaml
{"description": "Request body for creating content exports.", "type": "object", "properties": {"dlp_profile_id": {"description": "Filter by DLP profile IDs.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["e91a2360-da51-4fdf-9711-bcdecd462614"]}, "dlp_profile_information": {"description": "DLP profile metadata for the export.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_DLPProfile"}}, "integration_id": {"description": "Filter by integration IDs.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["c416bc38-75dc-425f-ae25-c37b5df5c37f"]}, "max_affliction_date": {"description": "Filter to view content flagged on or before this date.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}, "min_affliction_date": {"description": "Filter to view content flagged on or after this date.", "type": "string", "format": "date-time", "example": "2023-01-01T00:00:00Z"}, "orders": {"description": "Ordering specifications for the export.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_ContentOrder"}}, "search": {"description": "Search term to filter content.", "type": "string", "example": "sensitive"}, "vendors": {"description": "Filter by vendor types.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_VendorsEnum"}, "example": ["GOOGLE_WORKSPACE"]}}, "required": ["dlp_profile_information"]}
```
