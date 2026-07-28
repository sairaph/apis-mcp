---
title: List Apps
page_id: operation-get-accounts-account-id-magic-apps-fc1fb16b
path: operations/magic-account-apps
description: Lists Apps associated with an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/apps
operation_ids:
    - magic-account-apps-list-apps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Apps

`GET /accounts/{account_id}/magic/apps`

Operation ID: `magic-account-apps-list-apps`

Lists Apps associated with an account.

## Definition

```yaml
{"operationId": "magic-account-apps-list-apps", "summary": "List Apps", "description": "Lists Apps associated with an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Apps response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_apps_collection_response"}}}}, "4XX": {"description": "List Apps response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Account Apps"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.apps", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
