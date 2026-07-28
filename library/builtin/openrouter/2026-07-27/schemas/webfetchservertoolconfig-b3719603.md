---
title: WebFetchServerToolConfig
page_id: schema-webfetchservertoolconfig-b3719603
path: schemas
description: Configuration for the openrouter:web_fetch server tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebFetchServerToolConfig

Configuration for the openrouter:web_fetch server tool

```yaml
{"description": "Configuration for the openrouter:web_fetch server tool", "example": {"max_content_tokens": 100000, "max_uses": 10}, "properties": {"allowed_domains": {"description": "Only fetch from these domains.", "items": {"type": "string"}, "type": "array"}, "blocked_domains": {"description": "Never fetch from these domains.", "items": {"type": "string"}, "type": "array"}, "engine": {"$ref": "#/components/schemas/WebFetchEngineEnum"}, "max_content_tokens": {"description": "Maximum content length in approximate tokens. Content exceeding this limit is truncated.", "example": 100000, "type": "integer"}, "max_uses": {"description": "Maximum number of web fetches per request. Once exceeded, the tool returns an error.", "example": 10, "type": "integer"}}, "type": "object"}
```
