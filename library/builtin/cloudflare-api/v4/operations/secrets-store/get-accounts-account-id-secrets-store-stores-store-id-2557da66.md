---
title: Get a store by ID
page_id: operation-get-accounts-account-id-secrets-store-stores-store-id-79bf93d4
path: operations/secrets-store
description: Returns details of a single store.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}
operation_ids:
    - secrets-store-get-store-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a store by ID

`GET /accounts/{account_id}/secrets_store/stores/{store_id}`

Operation ID: `secrets-store-get-store-by-id`

Returns details of a single store.

## Definition

```yaml
{"operationId": "secrets-store-get-store-by-id", "summary": "Get a store by ID", "description": "Returns details of a single store.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}], "responses": {"200": {"description": "Store details.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_store_response"}}}}, "4XX": {"description": "Get store response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write", "Secrets Store Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.system-stores.get.by", "x-fern-sdk-method-name": "id", "x-forge-hidden": true, "x-stability": "beta"}
```
