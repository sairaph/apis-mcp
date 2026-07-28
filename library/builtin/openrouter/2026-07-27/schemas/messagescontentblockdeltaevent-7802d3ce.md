---
title: MessagesContentBlockDeltaEvent
page_id: schema-messagescontentblockdeltaevent-7802d3ce
path: schemas
description: Event sent when content is added to a content block
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesContentBlockDeltaEvent

Event sent when content is added to a content block

```yaml
{"description": "Event sent when content is added to a content block", "example": {"delta": {"text": "Hello", "type": "text_delta"}, "index": 0, "type": "content_block_delta"}, "properties": {"delta": {"oneOf": [{"properties": {"text": {"type": "string"}, "type": {"enum": ["text_delta"], "type": "string"}}, "required": ["type", "text"], "type": "object"}, {"properties": {"partial_json": {"type": "string"}, "type": {"enum": ["input_json_delta"], "type": "string"}}, "required": ["type", "partial_json"], "type": "object"}, {"properties": {"thinking": {"type": "string"}, "type": {"enum": ["thinking_delta"], "type": "string"}}, "required": ["type", "thinking"], "type": "object"}, {"properties": {"signature": {"type": "string"}, "type": {"enum": ["signature_delta"], "type": "string"}}, "required": ["type", "signature"], "type": "object"}, {"properties": {"citation": {"discriminator": {"mapping": {"char_location": "#/components/schemas/AnthropicCitationCharLocation", "content_block_location": "#/components/schemas/AnthropicCitationContentBlockLocation", "page_location": "#/components/schemas/AnthropicCitationPageLocation", "search_result_location": "#/components/schemas/AnthropicCitationSearchResultLocation", "web_search_result_location": "#/components/schemas/AnthropicCitationWebSearchResultLocation"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicCitationCharLocation"}, {"$ref": "#/components/schemas/AnthropicCitationPageLocation"}, {"$ref": "#/components/schemas/AnthropicCitationContentBlockLocation"}, {"$ref": "#/components/schemas/AnthropicCitationWebSearchResultLocation"}, {"$ref": "#/components/schemas/AnthropicCitationSearchResultLocation"}]}, "type": {"enum": ["citations_delta"], "type": "string"}}, "required": ["type", "citation"], "type": "object"}, {"properties": {"content": {"type": ["string", "null"]}, "encrypted_content": {"type": ["string", "null"]}, "type": {"enum": ["compaction_delta"], "type": "string"}}, "required": ["type", "content"], "type": "object"}]}, "index": {"type": "integer"}, "type": {"enum": ["content_block_delta"], "type": "string"}}, "required": ["type", "index", "delta"], "type": "object"}
```
