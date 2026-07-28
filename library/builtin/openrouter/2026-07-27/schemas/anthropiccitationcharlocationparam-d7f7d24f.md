---
title: AnthropicCitationCharLocationParam
page_id: schema-anthropiccitationcharlocationparam-d7f7d24f
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationCharLocationParam

```yaml
{"example": {"cited_text": "Example cited text", "document_index": 0, "document_title": null, "end_char_index": 18, "start_char_index": 0, "type": "char_location"}, "properties": {"cited_text": {"type": "string"}, "document_index": {"type": "integer"}, "document_title": {"type": ["string", "null"]}, "end_char_index": {"type": "integer"}, "start_char_index": {"type": "integer"}, "type": {"enum": ["char_location"], "type": "string"}}, "required": ["type", "cited_text", "document_index", "document_title", "start_char_index", "end_char_index"], "type": "object"}
```
