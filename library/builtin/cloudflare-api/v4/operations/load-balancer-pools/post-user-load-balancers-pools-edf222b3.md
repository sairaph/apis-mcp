---
title: Create Pool
page_id: operation-post-user-load-balancers-pools-b4d9f763
path: operations/load-balancer-pools
description: Create a new pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /user/load_balancers/pools
operation_ids:
    - load-balancer-pools-create-pool
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Pool

`POST /user/load_balancers/pools`

Operation ID: `load-balancer-pools-create-pool`

Create a new pool.

## Definition

```yaml
{"operationId": "load-balancer-pools-create-pool", "summary": "Create Pool", "description": "Create a new pool.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"check_regions": {"$ref": "#/components/schemas/load-balancing_check_regions"}, "description": {"$ref": "#/components/schemas/load-balancing_schemas-description"}, "enabled": {"$ref": "#/components/schemas/load-balancing_enabled"}, "latitude": {"$ref": "#/components/schemas/load-balancing_latitude"}, "load_shedding": {"$ref": "#/components/schemas/load-balancing_load_shedding"}, "longitude": {"$ref": "#/components/schemas/load-balancing_longitude"}, "minimum_origins": {"$ref": "#/components/schemas/load-balancing_minimum_origins"}, "monitor": {"$ref": "#/components/schemas/load-balancing_monitor_id"}, "monitor_group": {"$ref": "#/components/schemas/load-balancing_monitor_group_id"}, "name": {"$ref": "#/components/schemas/load-balancing_name"}, "networks": {"$ref": "#/components/schemas/load-balancing_networks"}, "notification_email": {"$ref": "#/components/schemas/load-balancing_notification_email"}, "notification_filter": {"$ref": "#/components/schemas/load-balancing_notification_filter"}, "origin_steering": {"$ref": "#/components/schemas/load-balancing_origin_steering"}, "origins": {"$ref": "#/components/schemas/load-balancing_origins"}}, "required": ["origins", "name"]}}}}, "responses": {"200": {"description": "Create Pool response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-single_response"}}}}, "4XX": {"description": "Create Pool response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
