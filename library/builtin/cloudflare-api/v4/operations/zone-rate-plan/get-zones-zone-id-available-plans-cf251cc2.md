---
title: List Available Plans
page_id: operation-get-zones-zone-id-available-plans-71575492
path: operations/zone-rate-plan
description: Lists available plans the zone can subscribe to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/available_plans
operation_ids:
    - zone-rate-plan-list-available-plans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Available Plans

`GET /zones/{zone_id}/available_plans`

Operation ID: `zone-rate-plan-list-available-plans`

Lists available plans the zone can subscribe to.

## Definition

```yaml
{"operationId": "zone-rate-plan-list-available-plans", "summary": "List Available Plans", "description": "Lists available plans the zone can subscribe to.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "List Available Plans response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/bill-subs-api_available-rate-plan"}}}}]}}}}, "4XX": {"description": "List Available Plans response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/bill-subs-api_available-rate-plan"}}}}]}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Rate Plan"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.plans", "x-fern-sdk-method-name": "list"}
```
