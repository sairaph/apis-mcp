---
title: posture-api_TuneFindingSeverityRequest
page_id: schema-posture-api-tunefindingseverityrequest-547c2c01
path: schemas
description: Request body for updating a finding's severity.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_TuneFindingSeverityRequest

Request body for updating a finding's severity.

```yaml
{"description": "Request body for updating a finding's severity.", "type": "object", "properties": {"new_severity": {"description": "The numeric severity value to apply to the finding.", "type": "integer", "enum": [1, 2, 3, 4]}}, "required": ["new_severity"]}
```
