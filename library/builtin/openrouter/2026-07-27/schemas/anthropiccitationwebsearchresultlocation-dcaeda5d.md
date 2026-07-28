---
title: AnthropicCitationWebSearchResultLocation
page_id: schema-anthropiccitationwebsearchresultlocation-dcaeda5d
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCitationWebSearchResultLocation

```yaml
{"example": {"cited_text": "Example cited text", "encrypted_index": "enc_idx_0", "title": "Example Page", "type": "web_search_result_location", "url": "https://example.com"}, "properties": {"cited_text": {"type": "string"}, "encrypted_index": {"type": "string"}, "title": {"type": ["string", "null"]}, "type": {"enum": ["web_search_result_location"], "type": "string"}, "url": {"type": "string"}}, "required": ["type", "cited_text", "encrypted_index", "title", "url"], "type": "object"}
```
