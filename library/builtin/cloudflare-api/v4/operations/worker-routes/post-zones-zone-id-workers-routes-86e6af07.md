---
title: Create Route
page_id: operation-post-zones-zone-id-workers-routes-f39bb238
path: operations/worker-routes
description: Creates a route that maps a URL pattern to a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/workers/routes
operation_ids:
    - worker-routes-create-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Route

`POST /zones/{zone_id}/workers/routes`

Operation ID: `worker-routes-create-route`

Creates a route that maps a URL pattern to a Worker.

## Definition

```yaml
{"operationId": "worker-routes-create-route", "summary": "Create Route", "description": "Creates a route that maps a URL pattern to a Worker.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_route"}}}}, "responses": {"200": {"description": "Create Route response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_route"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create Route response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Routes"], "x-api-token-group": ["Workers Routes Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.routes", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
