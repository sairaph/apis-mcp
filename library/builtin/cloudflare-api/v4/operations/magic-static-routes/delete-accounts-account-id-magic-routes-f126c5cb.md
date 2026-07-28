---
title: Delete Many Routes
page_id: operation-delete-accounts-account-id-magic-routes-11f56383
path: operations/magic-static-routes
description: Delete multiple Magic static routes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/routes
operation_ids:
    - magic-static-routes-delete-many-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Many Routes

`DELETE /accounts/{account_id}/magic/routes`

Operation ID: `magic-static-routes-delete-many-routes`

Delete multiple Magic static routes.

## Definition

```yaml
{"operationId": "magic-static-routes-delete-many-routes", "summary": "Delete Many Routes", "description": "Delete multiple Magic static routes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_route_delete_many_request"}}}}, "responses": {"200": {"description": "Delete Many Routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_multiple_route_delete_response"}}}}, "4XX": {"description": "Delete Many Routes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_multiple_route_delete_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Static Routes"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.routes", "x-fern-sdk-method-name": "empty", "x-forge-hidden": true}
```
