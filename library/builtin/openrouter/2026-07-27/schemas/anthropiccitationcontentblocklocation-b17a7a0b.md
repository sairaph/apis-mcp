---
title: AnthropicCitationContentBlockLocation
page_id: schema-anthropiccitationcontentblocklocation-b17a7a0b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationContentBlockLocation

```yaml
{"example": {"cited_text": "Example cited text", "document_index": 0, "document_title": null, "end_block_index": 1, "file_id": null, "start_block_index": 0, "type": "content_block_location"}, "properties": {"cited_text": {"type": "string"}, "document_index": {"type": "integer"}, "document_title": {"type": ["string", "null"]}, "end_block_index": {"type": "integer"}, "file_id": {"type": ["string", "null"]}, "start_block_index": {"type": "integer"}, "type": {"enum": ["content_block_location"], "type": "string"}}, "required": ["type", "cited_text", "document_index", "document_title", "start_block_index", "end_block_index", "file_id"], "type": "object"}
```
