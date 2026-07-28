---
title: Get Route
page_id: operation-get-zones-zone-id-workers-routes-route-id-3d6df51c
path: operations/worker-routes
description: Returns information about a route, including URL pattern and Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/workers/routes/{route_id}
operation_ids:
    - worker-routes-get-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Route

`GET /zones/{zone_id}/workers/routes/{route_id}`

Operation ID: `worker-routes-get-route`

Returns information about a route, including URL pattern and Worker.

## Definition

```yaml
{"operationId": "worker-routes-get-route", "summary": "Get Route", "description": "Returns information about a route, including URL pattern and Worker.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "Get Route response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_route"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get Route response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Routes"], "x-api-token-group": ["Workers Routes Write", "Workers Routes Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.routes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
