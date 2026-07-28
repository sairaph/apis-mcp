---
title: Create Worker Account Settings
page_id: operation-put-accounts-account-id-workers-account-settings-57602732
path: operations/worker-account-settings
description: Creates Worker account settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/account-settings
operation_ids:
    - worker-account-settings-create-worker-account-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Worker Account Settings

`PUT /accounts/{account_id}/workers/account-settings`

Operation ID: `worker-account-settings-create-worker-account-settings`

Creates Worker account settings for an account.

## Definition

```yaml
{"operationId": "worker-account-settings-create-worker-account-settings", "summary": "Create Worker Account Settings", "description": "Creates Worker account settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_account-settings"}}}}, "responses": {"200": {"description": "Create Worker Account Settings response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_account-settings"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create Worker Account Settings response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Account Settings"], "x-api-token-group": ["Account Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.account-settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
