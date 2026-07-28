---
title: RequestMetadata
page_id: schema-requestmetadata-6b4c8380
path: schemas
description: Metadata key-value pairs for the request. Keys must be ≤64 characters and cannot contain brackets. Values must be ≤512 characters. Maximum 16 pairs allowed.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RequestMetadata

Metadata key-value pairs for the request. Keys must be ≤64 characters and cannot contain brackets. Values must be ≤512 characters. Maximum 16 pairs allowed.

```yaml
{"additionalProperties": {"maxLength": 512, "type": "string"}, "description": "Metadata key-value pairs for the request. Keys must be ≤64 characters and cannot contain brackets. Values must be ≤512 characters. Maximum 16 pairs allowed.", "example": {"session_id": "abc-def-ghi", "user_id": "123"}, "type": ["object", "null"]}
```
