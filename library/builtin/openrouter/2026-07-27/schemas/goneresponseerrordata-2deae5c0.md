---
title: GoneResponseErrorData
page_id: schema-goneresponseerrordata-2deae5c0
path: schemas
description: Error data for GoneResponse
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GoneResponseErrorData

Error data for GoneResponse

```yaml
{"description": "Error data for GoneResponse", "example": {"code": 410, "message": "The Coinbase APIs used by this endpoint have been deprecated, so the Coinbase Commerce credits API has been removed. Use the web credits purchase flow instead."}, "properties": {"code": {"type": "integer"}, "message": {"type": "string"}, "metadata": {"additionalProperties": {}, "type": ["object", "null"]}}, "required": ["code", "message"], "type": "object"}
```
