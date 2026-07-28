---
title: aaa_silence_create_request
page_id: schema-aaa-silence-create-request-1784040d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_silence_create_request

```yaml
{"type": "object", "properties": {"end_time": {"description": "When the silence ends.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}, "policy_id": {"$ref": "#/components/schemas/aaa_policy-id"}, "start_time": {"description": "When the silence starts.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}}}
```
