---
title: Update Route
page_id: operation-put-accounts-account-id-magic-routes-route-id-182ee635
path: operations/magic-static-routes
description: Update a specific Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/routes/{route_id}
operation_ids:
    - magic-static-routes-update-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Route

`PUT /accounts/{account_id}/magic/routes/{route_id}`

Operation ID: `magic-static-routes-update-route`

Update a specific Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.

## Definition

```yaml
{"operationId": "magic-static-routes-update-route", "summary": "Update Route", "description": "Update a specific Magic static route. Use `?validate_only=true` as an optional query parameter to run validation only without persisting changes.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_update_request"}}}}, "responses": {"200": {"description": "Update Route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_modified_response"}}}}, "4XX": {"description": "Update Route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_route_modified_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
