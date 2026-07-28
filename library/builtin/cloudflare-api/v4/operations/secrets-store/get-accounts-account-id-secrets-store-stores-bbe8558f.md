---
title: List account stores
page_id: operation-get-accounts-account-id-secrets-store-stores-c082ed55
path: operations/secrets-store
description: Lists all the stores in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores
operation_ids:
    - secrets-store-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account stores

`GET /accounts/{account_id}/secrets_store/stores`

Operation ID: `secrets-store-list`

Lists all the stores in an account.

## Definition

```yaml
{"operationId": "secrets-store-list", "summary": "List account stores", "description": "Lists all the stores in an account.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_direction"}, {"$ref": "#/components/parameters/secrets-store_page"}, {"$ref": "#/components/parameters/secrets-store_per_page"}, {"$ref": "#/components/parameters/secrets-store_store_order"}], "responses": {"200": {"description": "List account stores response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_stores_response_collection"}}}}, "4XX": {"description": "List account stores response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write", "Secrets Store Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.stores", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
