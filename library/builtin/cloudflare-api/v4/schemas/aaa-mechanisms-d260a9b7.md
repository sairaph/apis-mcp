---
title: aaa_mechanisms
page_id: schema-aaa-mechanisms-d260a9b7
path: schemas
description: List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_mechanisms

List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.

```yaml
{"description": "List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.", "type": "object", "properties": {"email": {"type": "array", "items": {"properties": {"id": {"description": "The email address", "type": "string", "x-auditable": true}}, "type": "object"}}, "pagerduty": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/aaa_uuid"}}, "type": "object"}}, "webhooks": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/aaa_uuid"}}, "type": "object"}}}, "example": {"email": [{"id": "test@example.com"}], "pagerduty": [{"id": "e8133a15-00a4-4d69-aec1-32f70c51f6e5"}], "webhooks": [{"id": "14cc1190-5d2b-4b98-a696-c424cb2ad05f"}]}}
```
