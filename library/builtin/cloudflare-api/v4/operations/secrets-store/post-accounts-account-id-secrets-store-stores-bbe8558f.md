---
title: Create a store
page_id: operation-post-accounts-account-id-secrets-store-stores-41d3f1f8
path: operations/secrets-store
description: Creates a store in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores
operation_ids:
    - secrets-store-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a store

`POST /accounts/{account_id}/secrets_store/stores`

Operation ID: `secrets-store-create`

Creates a store in the account.

## Definition

```yaml
{"operationId": "secrets-store-create", "summary": "Create a store", "description": "Creates a store in the account.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_createStoreObject"}}}}, "responses": {"200": {"description": "Store details.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_store_response"}}}}, "4XX": {"description": "Create store response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.stores", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
