---
title: AnthropicUsage
page_id: schema-anthropicusage-a453b6eb
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicUsage

```yaml
{"example": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 100, "output_tokens": 50, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}, "properties": {"cache_creation": {"$ref": "#/components/schemas/AnthropicCacheCreation"}, "cache_creation_input_tokens": {"type": ["integer", "null"]}, "cache_read_input_tokens": {"type": ["integer", "null"]}, "inference_geo": {"type": ["string", "null"]}, "input_tokens": {"type": "integer"}, "output_tokens": {"type": "integer"}, "output_tokens_details": {"$ref": "#/components/schemas/AnthropicOutputTokensDetails"}, "server_tool_use": {"$ref": "#/components/schemas/AnthropicServerToolUsage"}, "service_tier": {"$ref": "#/components/schemas/AnthropicServiceTier"}}, "required": ["input_tokens", "output_tokens", "output_tokens_details", "cache_creation_input_tokens", "cache_read_input_tokens", "cache_creation", "inference_geo", "server_tool_use", "service_tier"], "type": "object"}
```
