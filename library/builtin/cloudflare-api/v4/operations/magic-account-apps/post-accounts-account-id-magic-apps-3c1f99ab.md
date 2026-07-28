---
title: Create a new App
page_id: operation-post-accounts-account-id-magic-apps-6510aeba
path: operations/magic-account-apps
description: Creates a new App for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/apps
operation_ids:
    - magic-account-apps-add-app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new App

`POST /accounts/{account_id}/magic/apps`

Operation ID: `magic-account-apps-add-app`

Creates a new App for an account

## Definition

```yaml
{"operationId": "magic-account-apps-add-app", "summary": "Create a new App", "description": "Creates a new App for an account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_add_single_request"}}}}, "responses": {"201": {"description": "Create Account App response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_single_response"}}}}, "4XX": {"description": "Create Account App response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Account Apps"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.apps", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
