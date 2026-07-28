---
title: MessagesErrorEvent
page_id: schema-messageserrorevent-0bf3d9cd
path: schemas
description: Error event in the stream
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesErrorEvent

Error event in the stream

```yaml
{"description": "Error event in the stream", "example": {"error": {"error_type": "provider_overloaded", "message": "Overloaded", "type": "overloaded_error"}, "type": "error"}, "properties": {"error": {"properties": {"error_type": {"$ref": "#/components/schemas/ApiErrorType"}, "message": {"type": "string"}, "type": {"type": "string"}}, "required": ["type", "message"], "type": "object"}, "openrouter_metadata": {"$ref": "#/components/schemas/OpenRouterMetadata"}, "type": {"enum": ["error"], "type": "string"}}, "required": ["type", "error"], "type": "object"}
```
