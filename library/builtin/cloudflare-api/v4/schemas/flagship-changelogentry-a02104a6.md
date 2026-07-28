---
title: flagship_ChangelogEntry
page_id: schema-flagship-changelogentry-a02104a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# flagship_ChangelogEntry

```yaml
{"oneOf": [{"properties": {"after": {"$ref": "#/components/schemas/flagship_Flag"}, "event": {"type": "string", "enum": ["create"]}, "flag_key": {"type": "string"}}, "required": ["flag_key", "event", "after"], "type": "object"}, {"properties": {"after": {"$ref": "#/components/schemas/flagship_Flag"}, "event": {"type": "string", "enum": ["delete"]}, "flag_key": {"type": "string"}}, "required": ["flag_key", "event", "after"], "type": "object"}, {"properties": {"after": {"$ref": "#/components/schemas/flagship_Flag"}, "diff": {"type": "object", "additionalProperties": {"properties": {"from": {"$ref": "#/components/schemas/flagship_JsonValue"}, "to": {"$ref": "#/components/schemas/flagship_JsonValue"}}, "type": "object"}}, "event": {"type": "string", "enum": ["update"]}, "flag_key": {"type": "string"}}, "required": ["flag_key", "event", "after", "diff"], "type": "object"}]}
```
