---
title: List Load Balancers
page_id: operation-get-zones-zone-id-load-balancers-2135a4c2
path: operations/load-balancers
description: List configured load balancers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/load_balancers
operation_ids:
    - load-balancers-list-load-balancers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Load Balancers

`GET /zones/{zone_id}/load_balancers`

Operation ID: `load-balancers-list-load-balancers`

List configured load balancers.

## Definition

```yaml
{"operationId": "load-balancers-list-load-balancers", "summary": "List Load Balancers", "description": "List configured load balancers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "responses": {"200": {"description": "List Load Balancers response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-response_collection"}}}}, "4XX": {"description": "List Load Balancers response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancers"], "x-api-token-group": ["Load Balancers Write", "Load Balancers Read"]}
```
