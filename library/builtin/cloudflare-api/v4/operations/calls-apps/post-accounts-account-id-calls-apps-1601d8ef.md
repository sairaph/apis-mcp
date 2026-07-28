---
title: Create a new app
page_id: operation-post-accounts-account-id-calls-apps-44a94f5f
path: operations/calls-apps
description: Creates a new Cloudflare calls app. An app is an unique enviroment where each Session can access all Tracks within the app.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/calls/apps
operation_ids:
    - calls-apps-create-a-new-app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new app

`POST /accounts/{account_id}/calls/apps`

Operation ID: `calls-apps-create-a-new-app`

Creates a new Cloudflare calls app. An app is an unique enviroment where each Session can access all Tracks within the app.

## Definition

```yaml
{"operationId": "calls-apps-create-a-new-app", "summary": "Create a new app", "description": "Creates a new Cloudflare calls app. An app is an unique enviroment where each Session can access all Tracks within the app.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_editable_fields"}}}}, "responses": {"201": {"description": "Created a new app", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_response_single_with_secret"}}}}}, "security": [{"api_token": []}], "tags": ["Calls Apps"], "x-api-token-group": ["Calls Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.sfu", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
