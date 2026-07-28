---
title: posture-api_DLPProfileEntry
page_id: schema-posture-api-dlpprofileentry-410abcbc
path: schemas
description: Entry within a DLP profile.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_DLPProfileEntry

Entry within a DLP profile.

```yaml
{"description": "Entry within a DLP profile.", "type": "object", "properties": {"id": {"description": "Unique identifier for the DLP profile entry.", "type": "string", "format": "uuid", "example": "55ba2c6c-8ef4-4b2e-9148-e75e8b6ccac1"}, "name": {"description": "Name of the DLP profile entry.", "type": "string", "example": "Credit Card Numbers"}, "profile_id": {"description": "ID of the parent DLP profile.", "type": "string", "format": "uuid", "example": "e91a2360-da51-4fdf-9711-bcdecd462614"}}, "required": ["id", "name", "profile_id"]}
```
