---
title: Delete Route
page_id: operation-delete-zones-zone-id-workers-routes-route-id-b12c5c2d
path: operations/worker-routes
description: Deletes a route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/workers/routes/{route_id}
operation_ids:
    - worker-routes-delete-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Route

`DELETE /zones/{zone_id}/workers/routes/{route_id}`

Operation ID: `worker-routes-delete-route`

Deletes a route.

## Definition

```yaml
{"operationId": "worker-routes-delete-route", "summary": "Delete Route", "description": "Deletes a route.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "Delete Route response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/workers_identifier"}}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete Route response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Routes"], "x-api-token-group": ["Workers Routes Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.routes", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
