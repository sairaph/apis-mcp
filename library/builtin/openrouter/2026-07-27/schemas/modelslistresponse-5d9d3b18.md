---
title: ModelsListResponse
page_id: schema-modelslistresponse-5d9d3b18
path: schemas
description: List of available models
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ModelsListResponse

List of available models

```yaml
{"description": "List of available models", "example": {"data": [{"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "canonical_slug": "openai/gpt-4", "context_length": 8192, "created": 1692901234, "default_parameters": null, "description": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "expiration_date": null, "id": "openai/gpt-4", "knowledge_cutoff": null, "links": {"details": "/api/v1/models/openai/gpt-4/endpoints"}, "name": "GPT-4", "per_request_limits": null, "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "supported_parameters": ["temperature", "top_p", "max_tokens", "frequency_penalty", "presence_penalty"], "supported_voices": null, "top_provider": {"context_length": 8192, "is_moderated": true, "max_completion_tokens": 4096}}], "links": {"next": "/api/v1/models?offset=500&limit=500"}, "total_count": 150}, "properties": {"data": {"$ref": "#/components/schemas/ModelsListResponseData"}, "links": {"description": "Pagination links", "properties": {"next": {"description": "URL for the next page of results, or null if this is the last page", "example": "/api/v1/models?offset=500&limit=500", "type": ["string", "null"]}}, "required": ["next"], "type": "object"}, "total_count": {"description": "Total number of models matching the query", "example": 150, "type": "integer"}}, "required": ["data", "total_count", "links"], "type": "object"}
```
