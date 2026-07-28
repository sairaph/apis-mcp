---
title: ActivityResponse
page_id: schema-activityresponse-f4670273
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ActivityResponse

```yaml
{"example": {"data": [{"byok_usage_inference": 0.012, "completion_tokens": 125, "date": "2025-08-24", "endpoint_id": "550e8400-e29b-41d4-a716-446655440000", "model": "openai/gpt-4.1", "model_permaslug": "openai/gpt-4.1-2025-04-14", "prompt_tokens": 50, "provider_name": "OpenAI", "reasoning_tokens": 25, "requests": 5, "usage": 0.015}]}, "properties": {"data": {"description": "List of activity items", "items": {"$ref": "#/components/schemas/ActivityItem"}, "type": "array"}}, "required": ["data"], "type": "object"}
```
