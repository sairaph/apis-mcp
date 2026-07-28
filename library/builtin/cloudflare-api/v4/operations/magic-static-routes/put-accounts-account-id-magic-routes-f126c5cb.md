---
title: Update Many Routes
page_id: operation-put-accounts-account-id-magic-routes-d9a7a368
path: operations/magic-static-routes
description: Update multiple Magic static routes. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes. Only fields for a route that need to be changed need be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/routes
operation_ids:
    - magic-static-routes-update-many-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Many Routes

`PUT /accounts/{account_id}/magic/routes`

Operation ID: `magic-static-routes-update-many-routes`

Update multiple Magic static routes. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes. Only fields for a route that need to be changed need be provided.

## Definition

```yaml
{"operationId": "magic-static-routes-update-many-routes", "summary": "Update Many Routes", "description": "Update multiple Magic static routes. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes. Only fields for a route that need to be changed need be provided.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_update_many_request"}}}}, "responses": {"200": {"description": "Update Many Routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_multiple_route_modified_response"}}}}, "4XX": {"description": "Update Many Routes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_multiple_route_modified_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "bulk-update", "x-forge-hidden": true}
```
