---
title: AnthropicCitationPageLocation
page_id: schema-anthropiccitationpagelocation-61213e33
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationPageLocation

```yaml
{"example": {"cited_text": "Example cited text", "document_index": 0, "document_title": null, "end_page_number": 2, "file_id": null, "start_page_number": 1, "type": "page_location"}, "properties": {"cited_text": {"type": "string"}, "document_index": {"type": "integer"}, "document_title": {"type": ["string", "null"]}, "end_page_number": {"type": "integer"}, "file_id": {"type": ["string", "null"]}, "start_page_number": {"type": "integer"}, "type": {"enum": ["page_location"], "type": "string"}}, "required": ["type", "cited_text", "document_index", "document_title", "start_page_number", "end_page_number", "file_id"], "type": "object"}
```
