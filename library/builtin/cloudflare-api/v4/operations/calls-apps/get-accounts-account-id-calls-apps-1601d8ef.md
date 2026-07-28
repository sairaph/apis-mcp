---
title: List apps
page_id: operation-get-accounts-account-id-calls-apps-b13ea4b3
path: operations/calls-apps
description: Lists all apps in the Cloudflare account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/calls/apps
operation_ids:
    - calls-apps-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List apps

`GET /accounts/{account_id}/calls/apps`

Operation ID: `calls-apps-list`

Lists all apps in the Cloudflare account

## Definition

```yaml
{"operationId": "calls-apps-list", "summary": "List apps", "description": "Lists all apps in the Cloudflare account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "List apps response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_response_collection"}}}}, "4XX": {"description": "List apps response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls Apps"], "x-api-token-group": ["Calls Write", "Calls Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.sfu", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
