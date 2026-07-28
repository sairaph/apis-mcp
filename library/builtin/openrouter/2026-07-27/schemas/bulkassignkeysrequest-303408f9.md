---
title: BulkAssignKeysRequest
page_id: schema-bulkassignkeysrequest-303408f9
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BulkAssignKeysRequest

```yaml
{"example": {"key_hashes": ["c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93"]}, "properties": {"key_hashes": {"description": "Array of API key hashes to assign to the guardrail", "example": ["c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93"], "items": {"minLength": 1, "type": "string"}, "minItems": 1, "type": "array"}}, "required": ["key_hashes"], "type": "object"}
```
