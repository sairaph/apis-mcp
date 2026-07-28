---
title: Delete a store
page_id: operation-delete-accounts-account-id-secrets-store-stores-store-id-c85f6d50
path: operations/secrets-store
description: |-
    Deletes a single store. By default, a store that still contains secrets
    cannot be deleted and returns HTTP 409 (Conflict) with the "store_not_empty"
    error. Pass `force=true` to cascade-delete all secrets in the store.
    Empty stores are always deleted regardless of the force parameter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/secrets_store/stores/{store_id}
operation_ids:
    - secrets-store-delete-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a store

`DELETE /accounts/{account_id}/secrets_store/stores/{store_id}`

Operation ID: `secrets-store-delete-by-id`

Deletes a single store. By default, a store that still contains secrets
cannot be deleted and returns HTTP 409 (Conflict) with the "store_not_empty"
error. Pass `force=true` to cascade-delete all secrets in the store.
Empty stores are always deleted regardless of the force parameter.

## Definition

```yaml
{"operationId": "secrets-store-delete-by-id", "summary": "Delete a store", "description": "Deletes a single store. By default, a store that still contains secrets\ncannot be deleted and returns HTTP 409 (Conflict) with the \"store_not_empty\"\nerror. Pass `force=true` to cascade-delete all secrets in the store.\nEmpty stores are always deleted regardless of the force parameter.\n", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}, {"$ref": "#/components/parameters/secrets-store_store_id"}, {"$ref": "#/components/parameters/secrets-store_force"}], "responses": {"200": {"description": "Store deleted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_delete_response"}}}}, "409": {"description": "Store is not empty. Use force=true to cascade-delete all secrets.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}, "4XX": {"description": "Delete store failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.stores", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
