---
title: Delete Account App
page_id: operation-delete-accounts-account-id-magic-apps-account-app-id-f707a128
path: operations/magic-account-apps
description: Deletes specific Account App.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/apps/{account_app_id}
operation_ids:
    - magic-account-apps-delete-app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Account App

`DELETE /accounts/{account_id}/magic/apps/{account_app_id}`

Operation ID: `magic-account-apps-delete-app`

Deletes specific Account App.

## Definition

```yaml
{"operationId": "magic-account-apps-delete-app", "summary": "Delete Account App", "description": "Deletes specific Account App.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete App response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_app_single_response"}}}}, "4XX": {"description": "Delete App response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Account Apps"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.apps", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
