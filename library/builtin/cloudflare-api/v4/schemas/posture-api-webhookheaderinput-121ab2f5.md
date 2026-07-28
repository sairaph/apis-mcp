---
title: posture-api_WebhookHeaderInput
page_id: schema-posture-api-webhookheaderinput-121ab2f5
path: schemas
description: A header to include in webhook requests. On Create/Evaluate, both key and value are required. On Update, omitting value means "keep existing value".
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_WebhookHeaderInput

A header to include in webhook requests. On Create/Evaluate, both key and value are required. On Update, omitting value means "keep existing value".

```yaml
{"description": "A header to include in webhook requests. On Create/Evaluate, both key and value are required. On Update, omitting value means \"keep existing value\".", "type": "object", "properties": {"key": {"description": "Header key name.", "type": "string", "example": "Authorization", "maxLength": 255}, "value": {"description": "Header value. Required on Create and Evaluate. On Update, omit or set to null to keep existing value.", "type": "string", "example": "Bearer token123", "maxLength": 4096, "nullable": true}}, "required": ["key"]}
```
