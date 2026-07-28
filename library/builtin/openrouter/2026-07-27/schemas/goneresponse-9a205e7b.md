---
title: GoneResponse
page_id: schema-goneresponse-9a205e7b
path: schemas
description: Gone - Endpoint has been permanently removed or deprecated
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GoneResponse

Gone - Endpoint has been permanently removed or deprecated

```yaml
{"description": "Gone - Endpoint has been permanently removed or deprecated", "example": {"error": {"code": 410, "message": "The Coinbase APIs used by this endpoint have been deprecated, so the Coinbase Commerce credits API has been removed. Use the web credits purchase flow instead."}}, "properties": {"error": {"$ref": "#/components/schemas/GoneResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
