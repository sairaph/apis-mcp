---
title: aaa_silence
page_id: schema-aaa-silence-43c1f2b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_silence

```yaml
{"type": "object", "properties": {"created_at": {"description": "When the silence was created.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}, "end_time": {"description": "When the silence ends.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}, "id": {"$ref": "#/components/schemas/aaa_silence-id"}, "policy_id": {"$ref": "#/components/schemas/aaa_policy-id"}, "start_time": {"description": "When the silence starts.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}, "updated_at": {"description": "When the silence was modified.", "type": "string", "example": "2022-01-01T00:00:00Z", "x-auditable": true}}}
```
