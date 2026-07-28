---
title: Delete Load Balancer
page_id: operation-delete-zones-zone-id-load-balancers-load-balancer-id-107e111c
path: operations/load-balancers
description: Delete a configured load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/load_balancers/{load_balancer_id}
operation_ids:
    - load-balancers-delete-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Load Balancer

`DELETE /zones/{zone_id}/load_balancers/{load_balancer_id}`

Operation ID: `load-balancers-delete-load-balancer`

Delete a configured load balancer.

## Definition

```yaml
{"operationId": "load-balancers-delete-load-balancer", "summary": "Delete Load Balancer", "description": "Delete a configured load balancer.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-id_response"}}}}, "4XX": {"description": "Delete Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_components-schemas-id_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancers"], "x-api-token-group": ["Load Balancers Write"]}
```
