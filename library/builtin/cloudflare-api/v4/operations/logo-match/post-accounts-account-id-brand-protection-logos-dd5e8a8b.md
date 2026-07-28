---
title: Create new saved logo queries from image files
page_id: operation-post-accounts-account-id-brand-protection-logos-11c9641b
path: operations/logo-match
description: Return new saved logo queries created from image files
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/brand-protection/logos
operation_ids:
    - postAccountsAccountIdBrandProtectionLogos
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new saved logo queries from image files

`POST /accounts/{account_id}/brand-protection/logos`

Operation ID: `postAccountsAccountIdBrandProtectionLogos`

Return new saved logo queries created from image files

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}]
```

## Definition

```yaml
{"operationId": "postAccountsAccountIdBrandProtectionLogos", "summary": "Create new saved logo queries from image files", "description": "Return new saved logo queries created from image files", "parameters": [{"name": "tag", "in": "query", "schema": {"type": "string"}}, {"name": "match_type", "in": "query", "schema": {"type": "string"}}, {"name": "threshold", "in": "query", "schema": {"type": "number"}}], "requestBody": {"required": true, "content": {"application/x-www-form-urlencoded": {"schema": {"$ref": "#/components/schemas/brand-protection-api_ImageFile"}}}}, "responses": {"201": {"description": "Created", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brand-protection-api_Logo"}}}}, "422": {"$ref": "#/components/responses/brand-protection-api_UNPROCESSABLE_CONTENT"}, "default": {"$ref": "#/components/responses/brand-protection-api_DEFAULT_ERROR"}}, "security": [{"api_token": []}], "tags": ["logo_match"], "x-api-token-group": ["Intel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logos", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
