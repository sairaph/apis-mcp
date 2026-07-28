---
title: Update Route
page_id: operation-put-zones-zone-id-workers-routes-route-id-0403d3cf
path: operations/worker-routes
description: Updates the URL pattern or Worker associated with a route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/workers/routes/{route_id}
operation_ids:
    - worker-routes-update-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Route

`PUT /zones/{zone_id}/workers/routes/{route_id}`

Operation ID: `worker-routes-update-route`

Updates the URL pattern or Worker associated with a route.

## Definition

```yaml
{"operationId": "worker-routes-update-route", "summary": "Update Route", "description": "Updates the URL pattern or Worker associated with a route.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_route"}}}}, "responses": {"200": {"description": "Update Route response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_route"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Update Route response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Routes"], "x-api-token-group": ["Workers Routes Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.routes", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
