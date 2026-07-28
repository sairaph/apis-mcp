---
title: UpdateBYOKKeyRequest
page_id: schema-updatebyokkeyrequest-3f8252a7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpdateBYOKKeyRequest

```yaml
{"example": {"disabled": false, "name": "Updated OpenAI Key"}, "properties": {"allowed_models": {"description": "Optional allowlist of model slugs this credential may be used for. `null` means no restriction.", "example": null, "items": {"type": "string"}, "maxItems": 100, "type": ["array", "null"]}, "allowed_user_ids": {"description": "Optional allowlist of user IDs that may use this credential. `null` means no restriction.", "example": null, "items": {"type": "string"}, "maxItems": 100, "type": ["array", "null"]}, "disabled": {"description": "Whether this credential is disabled.", "example": false, "type": "boolean"}, "is_fallback": {"description": "Whether this credential is treated as a fallback — used only after non-fallback keys for the same provider have been tried.", "example": false, "type": "boolean"}, "key": {"description": "A new raw provider API key to rotate the credential in-place. The previous key material is overwritten and the masked label is regenerated. Encrypted at rest and never returned in API responses.", "example": "sk-proj-newkey456...", "minLength": 1, "type": "string"}, "name": {"description": "Optional human-readable name for the credential.", "example": "Updated OpenAI Key", "maxLength": 255, "type": ["string", "null"]}}, "type": "object"}
```
