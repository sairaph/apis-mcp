---
title: dlp_EntryUpdateType
page_id: schema-dlp-entryupdatetype-93eceb2a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EntryUpdateType

```yaml
{"oneOf": [{"allOf": [{"$ref": "#/components/schemas/dlp_CustomEntryUpdateType"}, {"properties": {"type": {"type": "string", "enum": ["custom"]}}, "required": ["type"], "type": "object"}]}, {"properties": {"type": {"type": "string", "enum": ["predefined"]}}, "required": ["type"], "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["integration"]}}, "required": ["type"], "type": "object"}]}
```
