---
title: List Monitors
page_id: operation-get-user-load-balancers-monitors-0aae6383
path: operations/load-balancer-monitors
description: List configured monitors for a user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/monitors
operation_ids:
    - load-balancer-monitors-list-monitors
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Monitors

`GET /user/load_balancers/monitors`

Operation ID: `load-balancer-monitors-list-monitors`

List configured monitors for a user.

## Definition

```yaml
{"operationId": "load-balancer-monitors-list-monitors", "summary": "List Monitors", "description": "List configured monitors for a user.", "responses": {"200": {"description": "Successful list monitors response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-collection"}}}}, "4XX": {"description": "Failed list monitors response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
