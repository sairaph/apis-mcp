---
title: AnthropicWebFetchBlock
page_id: schema-anthropicwebfetchblock-25ee2c12
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebFetchBlock

```yaml
{"example": {"content": {"citations": null, "source": {"data": "", "media_type": "text/plain", "type": "text"}, "title": null, "type": "document"}, "retrieved_at": null, "type": "web_fetch_result", "url": "https://example.com"}, "properties": {"content": {"$ref": "#/components/schemas/AnthropicDocumentBlock"}, "retrieved_at": {"type": ["string", "null"]}, "type": {"enum": ["web_fetch_result"], "type": "string"}, "url": {"type": "string"}}, "required": ["content", "retrieved_at", "type", "url"], "type": "object"}
```
