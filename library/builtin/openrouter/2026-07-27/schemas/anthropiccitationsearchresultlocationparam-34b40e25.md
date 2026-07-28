---
title: AnthropicCitationSearchResultLocationParam
page_id: schema-anthropiccitationsearchresultlocationparam-34b40e25
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationSearchResultLocationParam

```yaml
{"example": {"cited_text": "Example cited text", "end_block_index": 1, "search_result_index": 0, "source": "example_source", "start_block_index": 0, "title": "Example Result", "type": "search_result_location"}, "properties": {"cited_text": {"type": "string"}, "end_block_index": {"type": "integer"}, "search_result_index": {"type": "integer"}, "source": {"type": "string"}, "start_block_index": {"type": "integer"}, "title": {"type": ["string", "null"]}, "type": {"enum": ["search_result_location"], "type": "string"}}, "required": ["type", "cited_text", "search_result_index", "source", "title", "start_block_index", "end_block_index"], "type": "object"}
```
