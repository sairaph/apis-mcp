---
title: dlp_PredefinedEntry
page_id: schema-dlp-predefinedentry-062d9430
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PredefinedEntry

```yaml
{"type": "object", "properties": {"confidence": {"$ref": "#/components/schemas/dlp_EntryConfidence"}, "enabled": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "profile_id": {"type": "string", "format": "uuid", "deprecated": true, "nullable": true}, "variant": {"allOf": [{"$ref": "#/components/schemas/dlp_PredefinedEntryVariant"}]}}, "required": ["id", "name", "enabled", "confidence"]}
```
