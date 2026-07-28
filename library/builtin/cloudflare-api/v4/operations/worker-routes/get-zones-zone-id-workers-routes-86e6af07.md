---
title: List Routes
page_id: operation-get-zones-zone-id-workers-routes-fb7e1e73
path: operations/worker-routes
description: Returns routes for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/workers/routes
operation_ids:
    - worker-routes-list-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Routes

`GET /zones/{zone_id}/workers/routes`

Operation ID: `worker-routes-list-routes`

Returns routes for a zone.

## Definition

```yaml
{"operationId": "worker-routes-list-routes", "summary": "List Routes", "description": "Returns routes for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "List Routes response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_route"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List Routes response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Routes"], "x-api-token-group": ["Workers Routes Write", "Workers Routes Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.routes", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
