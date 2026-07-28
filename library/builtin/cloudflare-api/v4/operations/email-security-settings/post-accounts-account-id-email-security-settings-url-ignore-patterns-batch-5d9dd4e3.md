---
title: Batch URL ignore patterns
page_id: operation-post-accounts-account-id-email-security-settings-url-ignore-patterns-bat-c12dbe0e
path: operations/email-security-settings
description: Sends a batch of URL ignore patterns API calls to execute together.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/url_ignore_patterns/batch
operation_ids:
    - email_security_batch_url_ignore_patterns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch URL ignore patterns

`POST /accounts/{account_id}/email-security/settings/url_ignore_patterns/batch`

Operation ID: `email_security_batch_url_ignore_patterns`

Sends a batch of URL ignore patterns API calls to execute together.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_batch_url_ignore_patterns", "summary": "Batch URL ignore patterns", "description": "Sends a batch of URL ignore patterns API calls to execute together.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}}, "required": ["id"], "type": "object"}, "maxItems": 1000}, "patches": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_UpdateUrlIgnorePattern"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_CreateUrlIgnorePattern"}, "maxItems": 1000}, "puts": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_CreateUrlIgnorePattern"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_UrlIgnorePatternId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}}, "required": ["deletes", "patches", "puts", "posts"]}}}}, "responses": {"200": {"description": "Batch operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeletedUrlIgnorePattern"}}, "patches": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}, "puts": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_UrlIgnorePattern"}}}, "required": ["deletes", "patches", "puts", "posts"]}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
