---
title: ConflictResponse
page_id: schema-conflictresponse-e4c87a97
path: schemas
description: Conflict - Resource conflict or concurrent modification
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ConflictResponse

Conflict - Resource conflict or concurrent modification

```yaml
{"description": "Conflict - Resource conflict or concurrent modification", "example": {"error": {"code": 409, "message": "Resource conflict. Please try again later."}}, "properties": {"error": {"$ref": "#/components/schemas/ConflictResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
