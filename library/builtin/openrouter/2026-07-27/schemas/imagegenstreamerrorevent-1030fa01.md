---
title: ImageGenStreamErrorEvent
page_id: schema-imagegenstreamerrorevent-1030fa01
path: schemas
description: Emitted when streaming generation fails after the SSE response starts
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenStreamErrorEvent

Emitted when streaming generation fails after the SSE response starts

```yaml
{"description": "Emitted when streaming generation fails after the SSE response starts", "example": {"error": {"code": "upstream_error", "message": "The upstream provider returned an error", "param": null, "type": "provider_error"}, "type": "error"}, "properties": {"error": {"description": "Provider error details", "properties": {"code": {"description": "Provider error code, when supplied", "type": ["string", "null"]}, "message": {"description": "Provider error message", "type": "string"}, "param": {"description": "Request parameter associated with the error, when supplied", "type": ["string", "null"]}, "type": {"description": "Provider error type, when supplied", "type": ["string", "null"]}}, "required": ["message"], "type": "object"}, "type": {"description": "The event type", "enum": ["error"], "type": "string"}}, "required": ["type", "error"], "type": "object"}
```
