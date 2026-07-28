---
title: AnthropicCitationCharLocation
page_id: schema-anthropiccitationcharlocation-ba34d267
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationCharLocation

```yaml
{"example": {"cited_text": "Example cited text", "document_index": 0, "document_title": null, "end_char_index": 18, "file_id": null, "start_char_index": 0, "type": "char_location"}, "properties": {"cited_text": {"type": "string"}, "document_index": {"type": "integer"}, "document_title": {"type": ["string", "null"]}, "end_char_index": {"type": "integer"}, "file_id": {"type": ["string", "null"]}, "start_char_index": {"type": "integer"}, "type": {"enum": ["char_location"], "type": "string"}}, "required": ["type", "cited_text", "document_index", "document_title", "start_char_index", "end_char_index", "file_id"], "type": "object"}
```
