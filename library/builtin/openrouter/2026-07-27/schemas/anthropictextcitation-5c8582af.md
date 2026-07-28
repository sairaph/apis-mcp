---
title: AnthropicTextCitation
page_id: schema-anthropictextcitation-5c8582af
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextCitation

```yaml
{"discriminator": {"mapping": {"char_location": "#/components/schemas/AnthropicCitationCharLocation", "content_block_location": "#/components/schemas/AnthropicCitationContentBlockLocation", "page_location": "#/components/schemas/AnthropicCitationPageLocation", "search_result_location": "#/components/schemas/AnthropicCitationSearchResultLocation", "web_search_result_location": "#/components/schemas/AnthropicCitationWebSearchResultLocation"}, "propertyName": "type"}, "example": {"cited_text": "Example text", "document_index": 0, "document_title": null, "end_char_index": 10, "file_id": null, "start_char_index": 0, "type": "char_location"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicCitationCharLocation"}, {"$ref": "#/components/schemas/AnthropicCitationPageLocation"}, {"$ref": "#/components/schemas/AnthropicCitationContentBlockLocation"}, {"$ref": "#/components/schemas/AnthropicCitationWebSearchResultLocation"}, {"$ref": "#/components/schemas/AnthropicCitationSearchResultLocation"}]}
```
