---
title: Duplicate Secret
page_id: operation-post-accounts-account-id-secrets-store-stores-store-id-secrets-secret-id-3af94b57
path: operations/secrets-store
description: Creates a duplicate of the secret, keeping the value.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}/duplicate
operation_ids:
    - secrets-store-duplicate-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Duplicate Secret

`POST /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}/duplicate`

Operation ID: `secrets-store-duplicate-by-id`

Creates a duplicate of the secret, keeping the value.

## Definition

```yaml
{"operationId": "secrets-store-duplicate-by-id", "summary": "Duplicate Secret", "description": "Creates a duplicate of the secret, keeping the value.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}, {"$ref": "#/components/parameters/secrets-store_secret_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_duplicateSecretObject"}}}}, "responses": {"200": {"description": "Secret detail.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_secret_response"}}}}, "4XX": {"description": "Duplicate secret response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets", "x-fern-sdk-method-name": "duplicate", "x-forge-hidden": true, "x-stability": "beta"}
```
