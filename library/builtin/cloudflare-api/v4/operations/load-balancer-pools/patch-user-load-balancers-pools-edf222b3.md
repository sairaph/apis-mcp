---
title: Patch Pools
page_id: operation-patch-user-load-balancers-pools-355eb5c3
path: operations/load-balancer-pools
description: Apply changes to a number of existing pools, overwriting the supplied properties. Pools are ordered by ascending `name`. Returns the list of affected pools. Supports the standard pagination query parameters, either `limit`/`offset` or `per_page`/`page`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /user/load_balancers/pools
operation_ids:
    - load-balancer-pools-patch-pools
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Pools

`PATCH /user/load_balancers/pools`

Operation ID: `load-balancer-pools-patch-pools`

Apply changes to a number of existing pools, overwriting the supplied properties. Pools are ordered by ascending `name`. Returns the list of affected pools. Supports the standard pagination query parameters, either `limit`/`offset` or `per_page`/`page`.

## Definition

```yaml
{"operationId": "load-balancer-pools-patch-pools", "summary": "Patch Pools", "description": "Apply changes to a number of existing pools, overwriting the supplied properties. Pools are ordered by ascending `name`. Returns the list of affected pools. Supports the standard pagination query parameters, either `limit`/`offset` or `per_page`/`page`.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "string", "properties": {"notification_email": {"$ref": "#/components/schemas/load-balancing_patch_pools_notification_email"}}}}}}, "responses": {"200": {"description": "Patch Pools response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-response_collection"}}}}, "4XX": {"description": "Patch Pools response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
