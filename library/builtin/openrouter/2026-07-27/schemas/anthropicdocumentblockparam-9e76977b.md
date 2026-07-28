---
title: AnthropicDocumentBlockParam
page_id: schema-anthropicdocumentblockparam-9e76977b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicDocumentBlockParam

```yaml
{"example": {"source": {"data": "Hello, world!", "media_type": "text/plain", "type": "text"}, "type": "document"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "citations": {"properties": {"enabled": {"type": "boolean"}}, "type": ["object", "null"]}, "context": {"type": ["string", "null"]}, "source": {"oneOf": [{"$ref": "#/components/schemas/AnthropicBase64PdfSource"}, {"$ref": "#/components/schemas/AnthropicPlainTextSource"}, {"properties": {"content": {"anyOf": [{"type": "string"}, {"items": {"discriminator": {"mapping": {"image": "#/components/schemas/AnthropicImageBlockParam", "text": "#/components/schemas/AnthropicTextBlockParam"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicTextBlockParam"}, {"$ref": "#/components/schemas/AnthropicImageBlockParam"}]}, "type": "array"}]}, "type": {"enum": ["content"], "type": "string"}}, "required": ["type", "content"], "type": "object"}, {"$ref": "#/components/schemas/AnthropicUrlPdfSource"}, {"$ref": "#/components/schemas/AnthropicFileDocumentSource"}]}, "title": {"type": ["string", "null"]}, "type": {"enum": ["document"], "type": "string"}}, "required": ["type", "source"], "type": "object"}
```
