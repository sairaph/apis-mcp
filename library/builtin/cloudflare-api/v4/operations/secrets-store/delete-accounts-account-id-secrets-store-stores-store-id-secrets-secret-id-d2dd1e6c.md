---
title: Delete a secret
page_id: operation-delete-accounts-account-id-secrets-store-stores-store-id-secrets-secret-00312a5a
path: operations/secrets-store
description: Deletes a single secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}
operation_ids:
    - secrets-store-secret-delete-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a secret

`DELETE /accounts/{account_id}/secrets_store/stores/{store_id}/secrets/{secret_id}`

Operation ID: `secrets-store-secret-delete-by-id`

Deletes a single secret.

## Definition

```yaml
{"operationId": "secrets-store-secret-delete-by-id", "summary": "Delete a secret", "description": "Deletes a single secret.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}, {"$ref": "#/components/parameters/secrets-store_secret_id"}], "responses": {"202": {"description": "Secret deletion accepted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_delete_response"}}}}, "4XX": {"description": "Delete secret failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
