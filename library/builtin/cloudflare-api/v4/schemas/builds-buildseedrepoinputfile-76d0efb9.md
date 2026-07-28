---
title: builds_BuildSeedRepoInputFile
page_id: schema-builds-buildseedrepoinputfile-76d0efb9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_BuildSeedRepoInputFile

```yaml
{"type": "object", "properties": {"content": {"example": "name = \"my-worker\"\nmain = \"src/index.js\"\n", "type": "string"}, "filename": {"type": "string", "example": "wrangler.toml"}, "isBase64": {"type": "boolean", "default": false}, "replace": {"description": "Text to replace in the file", "type": "string"}}, "required": ["filename", "content"]}
```
