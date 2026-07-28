---
title: Delete a token
page_id: operation-delete-accounts-account-id-ai-search-tokens-id-9273c01c
path: operations/ai-search-tokens
description: Permanently delete a stored AI Search credential. Credentials in use by an instance cannot be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-search/tokens/{id}
operation_ids:
    - ai-search-delete-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a token

`DELETE /accounts/{account_id}/ai-search/tokens/{id}`

Operation ID: `ai-search-delete-tokens`

Permanently delete a stored AI Search credential. Credentials in use by an instance cannot be deleted.

## Definition

```yaml
{"operationId": "ai-search-delete-tokens", "summary": "Delete a token", "description": "Permanently delete a stored AI Search credential. Credentials in use by an instance cannot be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "c3dc5f0b34a14ff8e1b3ec04895e1b22"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid", "example": "62af0db3-c410-40b2-9ee3-0e93f6dd1de0"}}], "responses": {"200": {"description": "Token deleted.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "additionalProperties": false}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}, "path": {"type": "array", "items": {"type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "404": {"description": "Token not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}, "409": {"description": "Token in use by instances.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Search Tokens"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai-search"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "ai-search.tokens", "x-fern-sdk-method-name": "delete", "x-forge-params": {"id": {"description": "Stored credential record ID.", "flagName": "token-id"}}, "x-forge-require-confirmation": "This operation permanently deletes the stored credential."}
```
