---
title: Patch a secret
page_id: operation-patch-accounts-account-id-secrets-store-stores-store-id-secrets-secret-i-83765829
path: operations/secrets-store
description: Updates a single secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}
operation_ids:
    - secrets-store-patch-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch a secret

`PATCH /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}`

Operation ID: `secrets-store-patch-by-id`

Updates a single secret.

## Definition

```yaml
{"operationId": "secrets-store-patch-by-id", "summary": "Patch a secret", "description": "Updates a single secret.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}, {"$ref": "#/components/parameters/secrets-store_secret_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_patchSecretObject"}}}}, "responses": {"200": {"description": "Secret detail.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_secret_response"}}}}, "4XX": {"description": "Patch secret response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-stability": "beta"}
```
