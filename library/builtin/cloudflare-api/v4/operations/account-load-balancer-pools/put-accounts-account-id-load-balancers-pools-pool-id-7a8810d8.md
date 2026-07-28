---
title: Update Pool
page_id: operation-put-accounts-account-id-load-balancers-pools-pool-id-2e99df60
path: operations/account-load-balancer-pools
description: Modify a configured pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/load_balancers/pools/{pool_id}
operation_ids:
    - account-load-balancer-pools-update-pool
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Pool

`PUT /accounts/{account_id}/load_balancers/pools/{pool_id}`

Operation ID: `account-load-balancer-pools-update-pool`

Modify a configured pool.

## Definition

```yaml
{"operationId": "account-load-balancer-pools-update-pool", "summary": "Update Pool", "description": "Modify a configured pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"check_regions": {"$ref": "#/components/schemas/load-balancing_check_regions"}, "description": {"$ref": "#/components/schemas/load-balancing_schemas-description"}, "disabled_at": {"$ref": "#/components/schemas/load-balancing_schemas-disabled_at"}, "enabled": {"$ref": "#/components/schemas/load-balancing_enabled"}, "latitude": {"$ref": "#/components/schemas/load-balancing_latitude"}, "load_shedding": {"$ref": "#/components/schemas/load-balancing_load_shedding"}, "longitude": {"$ref": "#/components/schemas/load-balancing_longitude"}, "minimum_origins": {"$ref": "#/components/schemas/load-balancing_minimum_origins"}, "monitor": {"$ref": "#/components/schemas/load-balancing_monitor_id"}, "monitor_group": {"$ref": "#/components/schemas/load-balancing_monitor_group_id"}, "name": {"$ref": "#/components/schemas/load-balancing_name"}, "notification_email": {"$ref": "#/components/schemas/load-balancing_notification_email"}, "notification_filter": {"$ref": "#/components/schemas/load-balancing_notification_filter"}, "origin_steering": {"$ref": "#/components/schemas/load-balancing_origin_steering"}, "origins": {"$ref": "#/components/schemas/load-balancing_origins"}}, "required": ["origins", "name"]}}}}, "responses": {"200": {"description": "Update Pool response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-single_response"}}}}, "4XX": {"description": "Update Pool response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
