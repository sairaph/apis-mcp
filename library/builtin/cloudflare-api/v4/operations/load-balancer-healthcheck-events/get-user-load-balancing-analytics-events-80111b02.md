---
title: List Healthcheck Events
page_id: operation-get-user-load-balancing-analytics-events-76157dc3
path: operations/load-balancer-healthcheck-events
description: List origin health changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancing_analytics/events
operation_ids:
    - load-balancer-healthcheck-events-list-healthcheck-events
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Healthcheck Events

`GET /user/load_balancing_analytics/events`

Operation ID: `load-balancer-healthcheck-events-list-healthcheck-events`

List origin health changes.

## Definition

```yaml
{"operationId": "load-balancer-healthcheck-events-list-healthcheck-events", "summary": "List Healthcheck Events", "description": "List origin health changes.", "parameters": [{"name": "until", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_until"}}, {"name": "pool_name", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_pool_name"}}, {"name": "origin_healthy", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_origin_healthy"}}, {"name": "pool_id", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "since", "in": "query", "schema": {"description": "Start date and time of requesting data period in the ISO8601 format.", "type": "string", "format": "date-time", "example": "2016-11-11T12:00:00Z"}}, {"name": "origin_name", "in": "query", "schema": {"description": "The name for the origin to filter.", "type": "string", "example": "primary-dc-1"}}, {"name": "pool_healthy", "in": "query", "schema": {"description": "If true, filter events where the pool status is healthy. If false, filter events where the pool status is unhealthy.", "type": "boolean", "example": true, "default": true}}], "responses": {"200": {"description": "List Healthcheck Events response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-response_collection"}}}}, "4XX": {"description": "List Healthcheck Events response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_components-schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Healthcheck Events"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
