---
title: ListWorkspacesResponse
page_id: schema-listworkspacesresponse-8614ec6a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListWorkspacesResponse

```yaml
{"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "created_by": "user_abc123", "default_guardrail_id": "595d5849-7e86-51fd-a7c0-705c34e4afff", "default_image_model": "openai/dall-e-3", "default_provider_sort": "price", "default_text_model": "openai/gpt-4o", "description": "Production environment workspace", "id": "550e8400-e29b-41d4-a716-446655440000", "io_logging_api_key_ids": null, "io_logging_sampling_rate": 1, "is_data_discount_logging_enabled": true, "is_observability_broadcast_enabled": false, "is_observability_io_logging_enabled": false, "name": "Production", "slug": "production", "updated_at": "2025-08-24T15:45:00Z"}], "total_count": 1}, "properties": {"data": {"description": "List of workspaces", "items": {"$ref": "#/components/schemas/Workspace"}, "type": "array"}, "total_count": {"description": "Total number of workspaces", "example": 5, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
