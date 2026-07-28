---
title: ActivityItem
page_id: schema-activityitem-3aeaa0a9
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ActivityItem

```yaml
{"example": {"byok_usage_inference": 0.012, "completion_tokens": 125, "date": "2025-08-24", "endpoint_id": "550e8400-e29b-41d4-a716-446655440000", "model": "openai/gpt-4.1", "model_permaslug": "openai/gpt-4.1-2025-04-14", "prompt_tokens": 50, "provider_name": "OpenAI", "reasoning_tokens": 25, "requests": 5, "usage": 0.015}, "properties": {"byok_usage_inference": {"description": "BYOK inference cost in USD (external credits spent)", "example": 0.012, "format": "double", "type": "number"}, "completion_tokens": {"description": "Total completion tokens generated", "example": 125, "type": "integer"}, "date": {"description": "Date of the activity (YYYY-MM-DD format)", "example": "2025-08-24", "type": "string"}, "endpoint_id": {"description": "Unique identifier for the endpoint", "example": "550e8400-e29b-41d4-a716-446655440000", "type": "string"}, "model": {"description": "Model slug (e.g., \"openai/gpt-4.1\")", "example": "openai/gpt-4.1", "type": "string"}, "model_permaslug": {"description": "Model permaslug (e.g., \"openai/gpt-4.1-2025-04-14\")", "example": "openai/gpt-4.1-2025-04-14", "type": "string"}, "prompt_tokens": {"description": "Total prompt tokens used", "example": 50, "type": "integer"}, "provider_name": {"description": "Name of the provider serving this endpoint", "example": "OpenAI", "type": "string"}, "reasoning_tokens": {"description": "Total reasoning tokens used", "example": 25, "type": "integer"}, "requests": {"description": "Number of requests made", "example": 5, "type": "integer"}, "usage": {"description": "Total cost in USD (OpenRouter credits spent)", "example": 0.015, "format": "double", "type": "number"}}, "required": ["date", "model", "model_permaslug", "endpoint_id", "provider_name", "usage", "byok_usage_inference", "requests", "prompt_tokens", "completion_tokens", "reasoning_tokens"], "type": "object"}
```
