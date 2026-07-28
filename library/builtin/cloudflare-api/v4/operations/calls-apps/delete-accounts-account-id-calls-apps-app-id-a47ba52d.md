---
title: Delete app
page_id: operation-delete-accounts-account-id-calls-apps-app-id-0cc8d8a8
path: operations/calls-apps
description: Deletes an app from Cloudflare Calls
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/calls/apps/{app_id}
operation_ids:
    - calls-apps-delete-app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete app

`DELETE /accounts/{account_id}/calls/apps/{app_id}`

Operation ID: `calls-apps-delete-app`

Deletes an app from Cloudflare Calls

## Definition

```yaml
{"operationId": "calls-apps-delete-app", "summary": "Delete app", "description": "Deletes an app from Cloudflare Calls", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "Delete app response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_response_single"}}}}, "4XX": {"description": "Delete app response failure", "content": {"application/json": {}}}}, "security": [{"api_token": []}], "tags": ["Calls Apps"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.sfu", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
