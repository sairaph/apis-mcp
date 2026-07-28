---
title: posture-api_FindingExportOrderFilterRequest
page_id: schema-posture-api-findingexportorderfilterrequest-09f35eeb
path: schemas
description: Order specification for finding exports.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingExportOrderFilterRequest

Order specification for finding exports.

```yaml
{"description": "Order specification for finding exports.", "type": "object", "properties": {"direction": {"$ref": "#/components/schemas/posture-api_DirectionEnum"}, "name": {"description": "Which field to use when ordering the findings.", "type": "string", "example": "instance_count", "enum": ["instance_count", "finding.name", "integration.name", "latest_affliction_date", "severity"]}}, "required": ["direction", "name"]}
```
