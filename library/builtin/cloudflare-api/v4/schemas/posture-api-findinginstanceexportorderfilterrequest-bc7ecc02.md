---
title: posture-api_FindingInstanceExportOrderFilterRequest
page_id: schema-posture-api-findinginstanceexportorderfilterrequest-bc7ecc02
path: schemas
description: Order specification for finding instance exports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingInstanceExportOrderFilterRequest

Order specification for finding instance exports.

```yaml
{"description": "Order specification for finding instance exports.", "type": "object", "properties": {"direction": {"$ref": "#/components/schemas/posture-api_DirectionEnum"}, "name": {"description": "Which field to use when ordering the finding instances.", "type": "string", "example": "asset.name", "enum": ["asset.name", "affliction_date"]}}, "required": ["direction", "name"]}
```
