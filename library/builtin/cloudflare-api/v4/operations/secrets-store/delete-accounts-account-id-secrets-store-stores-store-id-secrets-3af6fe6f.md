---
title: Delete secrets
page_id: operation-delete-accounts-account-id-secrets-store-stores-store-id-secrets-272347de
path: operations/secrets-store
description: Deletes one or more secrets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}/secrets
operation_ids:
    - secrets-store-delete-bulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete secrets

`DELETE /accounts/{account_id}/secrets_store/stores/{store_id}/secrets`

Operation ID: `secrets-store-delete-bulk`

Deletes one or more secrets.

## Definition

```yaml
{"operationId": "secrets-store-delete-bulk", "summary": "Delete secrets", "description": "Deletes one or more secrets.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_deleteSecretsRequest"}}}}, "responses": {"202": {"description": "Secrets deletion accepted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_delete_response"}}}}, "4XX": {"description": "Delete secrets response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.secrets.bulk", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
