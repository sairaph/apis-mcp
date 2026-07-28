---
title: AnthropicWebFetchContent
page_id: schema-anthropicwebfetchcontent-dacaf572
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicWebFetchContent

```yaml
{"discriminator": {"mapping": {"web_fetch_result": "#/components/schemas/AnthropicWebFetchBlock", "web_fetch_tool_result_error": "#/components/schemas/AnthropicWebFetchToolResultError"}, "propertyName": "type"}, "example": {"content": {"citations": null, "source": {"data": "", "media_type": "text/plain", "type": "text"}, "title": null, "type": "document"}, "retrieved_at": null, "type": "web_fetch_result", "url": "https://example.com"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicWebFetchToolResultError"}, {"$ref": "#/components/schemas/AnthropicWebFetchBlock"}]}
```
