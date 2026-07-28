---
title: Delete Route
page_id: operation-delete-accounts-account-id-magic-routes-route-id-abd06b56
path: operations/magic-static-routes
description: Disable and remove a specific Magic static route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/routes/{route_id}
operation_ids:
    - magic-static-routes-delete-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Route

`DELETE /accounts/{account_id}/magic/routes/{route_id}`

Operation ID: `magic-static-routes-delete-route`

Disable and remove a specific Magic static route.

## Definition

```yaml
{"operationId": "magic-static-routes-delete-route", "summary": "Delete Route", "description": "Disable and remove a specific Magic static route.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_deleted_response"}}}}, "4XX": {"description": "Delete Route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_route_deleted_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
