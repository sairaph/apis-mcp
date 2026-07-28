---
title: AppRankingsItem
page_id: schema-apprankingsitem-726faf5b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AppRankingsItem

```yaml
{"example": {"app_id": 12345, "app_name": "Cline", "rank": 1, "total_requests": 4321, "total_tokens": "12345678"}, "properties": {"app_id": {"description": "Stable numeric identifier of the app on OpenRouter.", "example": 12345, "type": "integer"}, "app_name": {"description": "Public display name of the app.", "example": "Cline", "type": "string"}, "rank": {"description": "1-based position of the app within this response, per the requested `sort`.", "example": 1, "type": "integer"}, "total_requests": {"description": "Number of requests attributed to the app inside the date window.", "example": 4321, "type": "integer"}, "total_tokens": {"description": "Sum of `prompt_tokens + completion_tokens` attributed to the app inside the date window, returned as a decimal string so 64-bit values are not truncated.", "example": "12345678", "type": "string"}}, "required": ["rank", "app_id", "app_name", "total_tokens", "total_requests"], "type": "object"}
```
