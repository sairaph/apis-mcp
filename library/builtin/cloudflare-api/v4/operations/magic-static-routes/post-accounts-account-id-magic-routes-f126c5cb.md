---
title: Create a Route
page_id: operation-post-accounts-account-id-magic-routes-c7af9945
path: operations/magic-static-routes
description: Creates a new Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/routes
operation_ids:
    - magic-static-routes-create-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Route

`POST /accounts/{account_id}/magic/routes`

Operation ID: `magic-static-routes-create-routes`

Creates a new Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.

## Definition

```yaml
{"operationId": "magic-static-routes-create-routes", "summary": "Create a Route", "description": "Creates a new Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_create_route_request"}}}}, "responses": {"200": {"description": "Create Routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_create_route_response"}}}}, "4XX": {"description": "Create Routes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_routes_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
