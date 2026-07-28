---
title: aaa_audit-log-action
page_id: schema-aaa-audit-log-action-f2033196
path: schemas
description: Provides information about the action performed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-log-action

Provides information about the action performed.

```yaml
{"description": "Provides information about the action performed.", "type": "object", "properties": {"description": {"description": "A short description of the action performed.", "type": "string", "example": "Add Member"}, "result": {"description": "The result of the action, indicating success or failure.", "type": "string", "example": "success"}, "time": {"description": "A timestamp indicating when the action was logged.", "type": "string", "format": "date-time", "example": "2024-04-26T17:31:07Z"}, "type": {"description": "A short string that describes the action that was performed.", "type": "string", "example": "create"}}}
```
