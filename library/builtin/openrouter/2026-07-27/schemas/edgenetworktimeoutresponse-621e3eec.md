---
title: EdgeNetworkTimeoutResponse
page_id: schema-edgenetworktimeoutresponse-621e3eec
path: schemas
description: Infrastructure Timeout - Provider request timed out at edge network
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# EdgeNetworkTimeoutResponse

Infrastructure Timeout - Provider request timed out at edge network

```yaml
{"description": "Infrastructure Timeout - Provider request timed out at edge network", "example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "properties": {"error": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
