---
title: Usage
page_id: schema-usage-8d59829c
path: schemas
description: Token usage information for the response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Usage

Token usage information for the response

```yaml
{"anyOf": [{"allOf": [{"$ref": "#/components/schemas/OpenAIResponsesUsage"}, {"properties": {"cost": {"description": "Cost of the completion", "format": "double", "type": ["number", "null"]}, "cost_details": {"properties": {"upstream_inference_cost": {"format": "double", "type": ["number", "null"]}, "upstream_inference_input_cost": {"format": "double", "type": "number"}, "upstream_inference_output_cost": {"format": "double", "type": "number"}}, "required": ["upstream_inference_input_cost", "upstream_inference_output_cost"], "type": "object"}, "is_byok": {"description": "Whether a request was made using a Bring Your Own Key configuration", "type": "boolean"}, "server_tool_use_details": {"$ref": "#/components/schemas/ServerToolUseDetails"}}, "type": "object"}]}, {"type": "null"}], "description": "Token usage information for the response", "example": {"cost": 0.0012, "cost_details": {"upstream_inference_cost": null, "upstream_inference_input_cost": 0.0008, "upstream_inference_output_cost": 0.0004}, "input_tokens": 10, "input_tokens_details": {"cached_tokens": 0}, "output_tokens": 25, "output_tokens_details": {"reasoning_tokens": 0}, "total_tokens": 35}}
```
