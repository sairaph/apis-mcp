---
title: posture-api_FindingInstanceExportFilterRequest
page_id: schema-posture-api-findinginstanceexportfilterrequest-d9895085
path: schemas
description: Filter specification for finding instance exports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingInstanceExportFilterRequest

Filter specification for finding instance exports.

```yaml
{"description": "Filter specification for finding instance exports.", "type": "object", "properties": {"archived": {"description": "Filter for archived status.", "type": "boolean", "example": false}, "max_affliction_date": {"description": "Filter to view findings that occurred on or before the affliction date. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}, "min_affliction_date": {"description": "Filter to view findings that occurred on or after the affliction date. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2023-01-01T00:00:00Z"}, "orders": {"description": "Ordering specifications for the export.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_FindingInstanceExportOrderFilterRequest"}, "default": []}, "search": {"description": "A search term.", "type": "string", "example": "sensitive data"}}}
```
