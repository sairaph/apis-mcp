---
title: Create a secret
page_id: operation-post-accounts-account-id-secrets-store-stores-store-id-secrets-2490a57c
path: operations/secrets-store
description: Creates a secret in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets
operation_ids:
    - secrets-store-secret-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a secret

`POST /accounts/{account_id}/secrets_store/stores/{store_id}/secrets`

Operation ID: `secrets-store-secret-create`

Creates a secret in the account.

## Definition

```yaml
{"operationId": "secrets-store-secret-create", "summary": "Create a secret", "description": "Creates a secret in the account.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/secrets-store_createSecretObject"}, "x-stainless-skip": ["terraform"]}}}}, "responses": {"200": {"description": "Secret detail.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_secrets_response_collection"}}}}, "4XX": {"description": "Create secret response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets.bulk", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
