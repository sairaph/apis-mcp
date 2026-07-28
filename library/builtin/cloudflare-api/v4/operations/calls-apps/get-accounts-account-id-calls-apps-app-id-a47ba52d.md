---
title: Retrieve app details
page_id: operation-get-accounts-account-id-calls-apps-app-id-a820b19b
path: operations/calls-apps
description: Fetches details for a single Calls app.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/calls/apps/{app_id}
operation_ids:
    - calls-apps-retrieve-app-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve app details

`GET /accounts/{account_id}/calls/apps/{app_id}`

Operation ID: `calls-apps-retrieve-app-details`

Fetches details for a single Calls app.

## Definition

```yaml
{"operationId": "calls-apps-retrieve-app-details", "summary": "Retrieve app details", "description": "Fetches details for a single Calls app.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/calls_account_identifier"}}], "responses": {"200": {"description": "Retrieve app details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_app_response_single"}}}}, "4XX": {"description": "Retrieve app details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/calls_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Calls Apps"], "x-api-token-group": ["Calls Write", "Calls Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "calls.sfu", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
