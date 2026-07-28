---
title: Route Details
page_id: operation-get-accounts-account-id-magic-routes-route-id-17a55a1d
path: operations/magic-static-routes
description: Get a specific Magic static route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/routes/{route_id}
operation_ids:
    - magic-static-routes-route-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Route Details

`GET /accounts/{account_id}/magic/routes/{route_id}`

Operation ID: `magic-static-routes-route-details`

Get a specific Magic static route.

## Definition

```yaml
{"operationId": "magic-static-routes-route-details", "summary": "Route Details", "description": "Get a specific Magic static route.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Route Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_single_response"}}}}, "4XX": {"description": "Route Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_route_single_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
