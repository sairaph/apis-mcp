---
title: moq_api-response-error
page_id: schema-moq-api-response-error-975dadbd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_api-response-error

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error codes:\n- 21001: Request body too large (HTTP 413)\n- 21002: Request body too small / missing name (HTTP 400)\n- 21003: Relay ID should be 32 hex characters (HTTP 400)\n- 21004: Failed to decode body — invalid JSON (HTTP 400)\n- 21005: Failed to read body (HTTP 400)\n- 21006: Unexpected server error (HTTP 500)\n- 21007: Relay not found (HTTP 404)\n- 21008: Relay limit exceeded for this account (HTTP 409)\n- 21009: Token limit reached — max 10 tokens per relay (HTTP 409)\n- 21010: Invalid operations — must be a non-empty subset of \"publish\" and \"subscribe\" (HTTP 400)\n- 21011: Invalid relay name — name must not be empty (HTTP 400)\n- 21012: Token expiry too long — expires must be no more than 1 year out (HTTP 400)\n- 21013: Invalid upstream URL — must be an absolute moqt:// or https:// URL with a host (HTTP 400)\n- 21014: Config cannot be set on create — set it via PUT after the relay exists (HTTP 400)\n", "type": "integer"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}, "example": [{"code": 21007, "message": "A MoQ relay with this ID was not found."}]}, "messages": {"type": "array", "items": {"type": "object"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "example": false}}, "required": ["success", "errors", "messages"]}
```
