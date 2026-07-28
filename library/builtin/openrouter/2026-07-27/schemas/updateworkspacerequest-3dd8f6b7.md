---
title: UpdateWorkspaceRequest
page_id: schema-updateworkspacerequest-3dd8f6b7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpdateWorkspaceRequest

```yaml
{"example": {"name": "Updated Workspace", "slug": "updated-workspace"}, "properties": {"default_image_model": {"description": "Default image model for this workspace", "example": "openai/dall-e-3", "type": ["string", "null"]}, "default_provider_sort": {"description": "Default provider sort preference (price, throughput, latency, exacto)", "example": "price", "type": ["string", "null"]}, "default_text_model": {"description": "Default text model for this workspace", "example": "openai/gpt-4o", "type": ["string", "null"]}, "description": {"description": "New description for the workspace", "example": "Updated description", "maxLength": 500, "type": ["string", "null"]}, "io_logging_api_key_ids": {"description": "Optional array of API key IDs to filter I/O logging", "example": null, "items": {"type": "integer"}, "type": ["array", "null"]}, "io_logging_sampling_rate": {"description": "Sampling rate for I/O logging (0.0001-1)", "example": 1, "format": "double", "type": "number"}, "is_data_discount_logging_enabled": {"description": "Whether data discount logging is enabled", "example": true, "type": "boolean"}, "is_observability_broadcast_enabled": {"description": "Whether broadcast is enabled", "example": false, "type": "boolean"}, "is_observability_io_logging_enabled": {"description": "Whether private logging is enabled", "example": false, "type": "boolean"}, "name": {"description": "New name for the workspace", "example": "Updated Workspace", "maxLength": 100, "minLength": 1, "type": "string"}, "slug": {"description": "New URL-friendly slug (lowercase alphanumeric segments separated by single hyphens, no leading/trailing hyphens)", "example": "updated-workspace", "maxLength": 50, "minLength": 1, "pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$", "type": "string"}}, "type": "object"}
```
