---
title: posture-api_DLPProfile
page_id: schema-posture-api-dlpprofile-8c8ec479
path: schemas
description: DLP profile configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_DLPProfile

DLP profile configuration.

```yaml
{"description": "DLP profile configuration.", "type": "object", "properties": {"entries": {"description": "Entries contained within this DLP profile.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_DLPProfileEntry"}}, "id": {"description": "Unique identifier for the DLP profile.", "type": "string", "format": "uuid", "example": "e91a2360-da51-4fdf-9711-bcdecd462614"}, "name": {"description": "Name of the DLP profile.", "type": "string", "example": "Financial Information"}}, "required": ["id", "name", "entries"]}
```
