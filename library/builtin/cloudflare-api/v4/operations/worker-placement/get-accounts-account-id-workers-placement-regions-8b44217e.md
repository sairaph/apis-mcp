---
title: List Placement Regions
page_id: operation-get-accounts-account-id-workers-placement-regions-a114d040
path: operations/worker-placement
description: Returns a list of available placement regions organized by cloud provider. These regions can be used to configure Smart Placement for Workers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/placement/regions
operation_ids:
    - worker-placement-list-regions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Placement Regions

`GET /accounts/{account_id}/workers/placement/regions`

Operation ID: `worker-placement-list-regions`

Returns a list of available placement regions organized by cloud provider. These regions can be used to configure Smart Placement for Workers.

## Definition

```yaml
{"operationId": "worker-placement-list-regions", "summary": "List Placement Regions", "description": "Returns a list of available placement regions organized by cloud provider. These regions can be used to configure Smart Placement for Workers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "List Placement Regions response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_placement-regions-response"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List Placement Regions response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Placement"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.placement.regions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
