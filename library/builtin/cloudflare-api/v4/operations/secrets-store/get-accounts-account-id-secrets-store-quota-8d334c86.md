---
title: View secret usage
page_id: operation-get-accounts-account-id-secrets-store-quota-be7c7500
path: operations/secrets-store
description: Lists the number of secrets used in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secrets_store/quota
operation_ids:
    - secrets-store-quota
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# View secret usage

`GET /accounts/{account_id}/secrets_store/quota`

Operation ID: `secrets-store-quota`

Lists the number of secrets used in the account.

## Definition

```yaml
{"operationId": "secrets-store-quota", "summary": "View secret usage", "description": "Lists the number of secrets used in the account.", "parameters": [{"$ref": "#/components/parameters/secrets-store_account_id"}], "responses": {"200": {"description": "Usage and quota.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secrets-store_quota_response"}}}}, "4XX": {"description": "View quota response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secrets-store_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Secrets Store"], "x-api-token-group": ["Secrets Store Write", "Secrets Store Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "secrets-store.quota", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
