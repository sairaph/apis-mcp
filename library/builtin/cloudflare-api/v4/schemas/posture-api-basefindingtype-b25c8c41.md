---
title: posture-api_BaseFindingType
page_id: schema-posture-api-basefindingtype-b25c8c41
path: schemas
description: Basic finding type information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_BaseFindingType

Basic finding type information.

```yaml
{"description": "Basic finding type information.", "type": "object", "properties": {"category": {"$ref": "#/components/schemas/posture-api_FindingCategory"}, "id": {"description": "The unique identifier of the finding.", "type": "string", "format": "uuid", "example": "a20895dd-9c3b-43bd-a608-71c98c6c2d94"}, "name": {"description": "The name of the finding.", "type": "string", "example": "Slack File Publicly Accessible"}, "severity": {"$ref": "#/components/schemas/posture-api_SeverityEnum"}, "vendor": {"description": "The SaaS/Cloud vendor of the platform with which the finding is associated.", "type": "string", "example": "Google Workspace"}}, "required": ["id", "name", "category", "vendor", "severity"]}
```
