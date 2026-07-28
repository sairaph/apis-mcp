---
title: FileParserPlugin
page_id: schema-fileparserplugin-6a200093
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FileParserPlugin

```yaml
{"example": {"enabled": true, "id": "file-parser", "pdf": {"engine": "cloudflare-ai"}}, "properties": {"enabled": {"description": "Set to false to disable the file-parser plugin for this request. Defaults to true.", "type": "boolean"}, "id": {"enum": ["file-parser"], "type": "string"}, "pdf": {"$ref": "#/components/schemas/PDFParserOptions"}}, "required": ["id"], "type": "object"}
```
