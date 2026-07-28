---
title: OutputWebFetchServerToolItem
page_id: schema-outputwebfetchservertoolitem-24b8aaa3
path: schemas
description: An openrouter:web_fetch server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputWebFetchServerToolItem

An openrouter:web_fetch server tool output item

```yaml
{"description": "An openrouter:web_fetch server tool output item", "example": {"httpStatus": 200, "id": "wf_tmp_abc123", "status": "completed", "title": "Example Domain", "type": "openrouter:web_fetch", "url": "https://example.com"}, "properties": {"content": {"type": "string"}, "error": {"description": "The error message if the fetch failed.", "type": "string"}, "httpStatus": {"description": "The HTTP status code returned by the upstream URL fetch.", "type": "integer"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "title": {"type": "string"}, "type": {"enum": ["openrouter:web_fetch"], "type": "string"}, "url": {"type": "string"}}, "required": ["status", "type"], "type": "object"}
```
