---
title: List store secrets
page_id: operation-get-accounts-account-id-secrets-store-stores-store-id-secrets-79eeb4cb
path: operations/secrets-store
description: Lists all store secrets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets
operation_ids:
    - secrets-store-secrets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List store secrets

`GET /accounts/{account_id}/secrets_store/stores/{store_id}/secrets`

Operation ID: `secrets-store-secrets-list`

Lists all store secrets.

## Definition

```yaml
{"operationId": "secrets-store-secrets-list", "summary": "List store secrets", "description": "Lists all store secrets.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}, {"$ref": "#/components/parameters/secrets-store_direction"}, {"$ref": "#/components/parameters/secrets-store_page"}, {"$ref": "#/components/parameters/secrets-store_per_page"}, {"$ref": "#/components/parameters/secrets-store_search"}, {"$ref": "#/components/parameters/secrets-store_order"}, {"$ref": "#/components/parameters/secrets-store_scopes_query"}], "responses": {"200": {"description": "List store secrets response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_secrets_response_collection"}}}}, "4XX": {"description": "List store secrets response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write", "Secrets Store Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-stability": "beta"}
```
