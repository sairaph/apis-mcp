---
title: posture-api_FindingBulkActionRequest
page_id: schema-posture-api-findingbulkactionrequest-9e7485eb
path: schemas
description: Request body for bulk actions on findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingBulkActionRequest

Request body for bulk actions on findings.

```yaml
{"description": "Request body for bulk actions on findings.", "type": "object", "properties": {"checks": {"description": "A list of finding IDs to pass along.", "type": "array", "items": {"maxLength": 512, "minLength": 1, "type": "string"}, "example": ["MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAxOjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMgo="]}}, "required": ["checks"]}
```
