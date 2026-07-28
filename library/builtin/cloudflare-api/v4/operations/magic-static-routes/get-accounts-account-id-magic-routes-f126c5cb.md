---
title: List Routes
page_id: operation-get-accounts-account-id-magic-routes-5c9d33fd
path: operations/magic-static-routes
description: List all Magic static routes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/routes
operation_ids:
    - magic-static-routes-list-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Routes

`GET /accounts/{account_id}/magic/routes`

Operation ID: `magic-static-routes-list-routes`

List all Magic static routes.

## Definition

```yaml
{"operationId": "magic-static-routes-list-routes", "summary": "List Routes", "description": "List all Magic static routes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_routes_collection_response"}}}}, "4XX": {"description": "List Routes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_routes_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
