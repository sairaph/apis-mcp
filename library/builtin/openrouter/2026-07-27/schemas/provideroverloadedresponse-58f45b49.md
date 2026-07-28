---
title: ProviderOverloadedResponse
page_id: schema-provideroverloadedresponse-58f45b49
path: schemas
description: Provider Overloaded - Provider is temporarily overloaded
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ProviderOverloadedResponse

Provider Overloaded - Provider is temporarily overloaded

```yaml
{"description": "Provider Overloaded - Provider is temporarily overloaded", "example": {"error": {"code": 529, "message": "Provider returned error"}}, "properties": {"error": {"$ref": "#/components/schemas/ProviderOverloadedResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
