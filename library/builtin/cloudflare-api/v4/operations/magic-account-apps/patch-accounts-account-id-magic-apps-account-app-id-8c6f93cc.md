---
title: Update an App
page_id: operation-patch-accounts-account-id-magic-apps-account-app-id-6792cf73
path: operations/magic-account-apps
description: Updates an Account App
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/apps/{account_app_id}
operation_ids:
    - magic-account-apps-patch-app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an App

`PATCH /accounts/{account_id}/magic/apps/{account_app_id}`

Operation ID: `magic-account-apps-patch-app`

Updates an Account App

## Definition

```yaml
{"operationId": "magic-account-apps-patch-app", "summary": "Update an App", "description": "Updates an Account App", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_update_request"}}}}, "responses": {"200": {"description": "Update App response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_single_response"}}}}, "4XX": {"description": "Update App response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Account Apps"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.apps", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
