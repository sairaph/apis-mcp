---
title: WebFetchPlugin
page_id: schema-webfetchplugin-88d4b3c6
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebFetchPlugin

```yaml
{"example": {"id": "web-fetch", "max_uses": 10}, "properties": {"allowed_domains": {"description": "Only fetch from these domains.", "items": {"type": "string"}, "type": "array"}, "blocked_domains": {"description": "Never fetch from these domains.", "items": {"type": "string"}, "type": "array"}, "id": {"enum": ["web-fetch"], "type": "string"}, "max_content_tokens": {"description": "Maximum content length in approximate tokens. Content exceeding this limit is truncated.", "type": "integer"}, "max_uses": {"description": "Maximum number of web fetches per request. Once exceeded, the tool returns an error.", "type": "integer"}}, "required": ["id"], "type": "object"}
```
