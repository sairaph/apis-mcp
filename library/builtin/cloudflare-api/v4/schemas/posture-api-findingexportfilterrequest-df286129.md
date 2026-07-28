---
title: posture-api_FindingExportFilterRequest
page_id: schema-posture-api-findingexportfilterrequest-df286129
path: schemas
description: Filter specification for findings export jobs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingExportFilterRequest

Filter specification for findings export jobs.

```yaml
{"description": "Filter specification for findings export jobs.", "type": "object", "properties": {"ignored": {"description": "Filter for only the ignored findings. Set to false to only see active items.", "type": "boolean", "example": true}, "integration_id": {"description": "Filter by multiple integration IDs.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["55d7337e-1d0a-49fc-9826-925ba40df035"]}, "max_affliction_date": {"description": "Filter to view findings that occurred on or before the affliction date. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2019-08-24T14:15:22Z"}, "min_affliction_date": {"description": "Filter to view findings that occurred on or after the affliction date. Can be a date-time in ISO 8601 format or an epoch timestamp.", "type": "string", "format": "date-time", "example": "2018-08-24T14:15:22Z"}, "orders": {"description": "Which fields to use when ordering the findings.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_FindingExportOrderFilterRequest"}}, "product": {"description": "Filter by finding's category product.", "type": "string", "example": "SaaS", "nullable": true, "oneOf": [{"$ref": "#/components/schemas/posture-api_ProductEnum"}, {"$ref": "#/components/schemas/posture-api_NullEnum"}]}, "search": {"description": "A search term.", "type": "string", "example": "public access"}, "severities": {"description": "Filter by severity levels.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_SeveritiesEnum"}, "example": ["CRITICAL"]}, "vendors": {"description": "Filter by vendor types.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_VendorsEnum"}, "example": ["AWS"]}}}
```
