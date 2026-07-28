---
title: aaa_audit-logs-v2
page_id: schema-aaa-audit-logs-v2-29c19a55
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2

```yaml
{"type": "object", "properties": {"account": {"description": "Contains account related information.", "type": "object", "properties": {"id": {"description": "A unique identifier for the account.", "type": "string", "example": "4bb334f7c94c4a29a045f03944f072e5"}, "name": {"description": "A string that identifies the account name.", "type": "string", "example": "Example Account"}}}, "action": {"$ref": "#/components/schemas/aaa_audit-log-action"}, "actor": {"allOf": [{"$ref": "#/components/schemas/aaa_audit-log-actor-base"}, {"properties": {"type": {"description": "The type of actor.", "type": "string", "example": "user", "enum": ["account", "cloudflare_admin", "system", "user"]}}, "type": "object"}]}, "id": {"$ref": "#/components/schemas/aaa_identifier-2"}, "raw": {"$ref": "#/components/schemas/aaa_audit-log-raw"}, "resource": {"description": "Provides details about the affected resource.", "type": "object", "properties": {"id": {"description": "The unique identifier for the affected resource.", "type": "string"}, "product": {"description": "The Cloudflare product associated with the resource.", "type": "string", "example": "members"}, "request": {"type": "object"}, "response": {"type": "object"}, "scope": {"description": "The scope of the resource.", "type": "object", "example": "accounts"}, "type": {"description": "The type of the resource.", "type": "string"}}}, "zone": {"description": "Provides details about the zone affected by the action.", "type": "object", "properties": {"id": {"description": "A string that identifies the zone id.", "type": "string"}, "name": {"description": "A string that identifies the zone name.", "type": "string", "example": "example.com"}}}}}
```
