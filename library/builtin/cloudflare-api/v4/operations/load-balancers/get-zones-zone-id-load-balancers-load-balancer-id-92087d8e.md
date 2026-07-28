---
title: Load Balancer Details
page_id: operation-get-zones-zone-id-load-balancers-load-balancer-id-c9410737
path: operations/load-balancers
description: Fetch a single configured load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/load_balancers/{load_balancer_id}
operation_ids:
    - load-balancers-load-balancer-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Load Balancer Details

`GET /zones/{zone_id}/load_balancers/{load_balancer_id}`

Operation ID: `load-balancers-load-balancer-details`

Fetch a single configured load balancer.

## Definition

```yaml
{"operationId": "load-balancers-load-balancer-details", "summary": "Load Balancer Details", "description": "Fetch a single configured load balancer.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "responses": {"200": {"description": "Load Balancer Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Load Balancer Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancers"], "x-api-token-group": ["Load Balancers Write", "Load Balancers Read"]}
```
