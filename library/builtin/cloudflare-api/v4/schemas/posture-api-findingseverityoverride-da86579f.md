---
title: posture-api_FindingSeverityOverride
page_id: schema-posture-api-findingseverityoverride-da86579f
path: schemas
description: Override information for finding severity.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingSeverityOverride

Override information for finding severity.

```yaml
{"description": "Override information for finding severity.", "type": "object", "properties": {"created_by": {"description": "User ID who created the override.", "type": "string", "example": "1234"}, "severity": {"$ref": "#/components/schemas/posture-api_SeverityEnum"}}, "example": {"created_by": "1234", "severity": "Critical"}, "required": ["created_by", "severity"]}
```
