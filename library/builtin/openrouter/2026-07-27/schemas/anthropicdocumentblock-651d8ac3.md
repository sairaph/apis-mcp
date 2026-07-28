---
title: AnthropicDocumentBlock
page_id: schema-anthropicdocumentblock-651d8ac3
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicDocumentBlock

```yaml
{"example": {"citations": null, "source": {"data": "Hello, world!", "media_type": "text/plain", "type": "text"}, "title": null, "type": "document"}, "properties": {"citations": {"$ref": "#/components/schemas/AnthropicCitationsConfig"}, "source": {"anyOf": [{"$ref": "#/components/schemas/AnthropicBase64PdfSource"}, {"$ref": "#/components/schemas/AnthropicPlainTextSource"}]}, "title": {"type": ["string", "null"]}, "type": {"enum": ["document"], "type": "string"}}, "required": ["source", "title", "type"], "type": "object"}
```
