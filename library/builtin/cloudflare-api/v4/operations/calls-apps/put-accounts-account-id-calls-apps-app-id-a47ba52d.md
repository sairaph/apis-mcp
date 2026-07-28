---
title: Edit app details
page_id: operation-put-accounts-account-id-calls-apps-app-id-27ac79ac
path: operations/calls-apps
description: Edit details for a single app.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/calls/apps/{app_id}
operation_ids:
    - calls-apps-update-app-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit app details

`PUT /accounts/{account_id}/calls/apps/{app_id}`

Operation ID: `calls-apps-update-app-details`

Edit details for a single app.

## Definition

```yaml
{"operationId": "calls-apps-update-app-details", "summary": "Edit app details", "description": "Edit details for a single app.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_editable_fields"}}}}, "responses": {"200": {"description": "Edit app details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_response_single"}}}}, "4XX": {"description": "Edit app details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls Apps"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.sfu", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
