---
title: List Available Rate Plans
page_id: operation-get-zones-zone-id-available-rate-plans-56298192
path: operations/zone-rate-plan
description: Lists all rate plans the zone can subscribe to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/available_rate_plans
operation_ids:
    - zone-rate-plan-list-available-rate-plans
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Available Rate Plans

`GET /zones/{zone_id}/available_rate_plans`

Operation ID: `zone-rate-plan-list-available-rate-plans`

Lists all rate plans the zone can subscribe to.

## Definition

```yaml
{"operationId": "zone-rate-plan-list-available-rate-plans", "summary": "List Available Rate Plans", "description": "Lists all rate plans the zone can subscribe to.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "List Available Rate Plans response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_plan_response_collection"}}}}, "4XX": {"description": "List Available Rate Plans response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_plan_response_collection"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Rate Plan"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.rate-plans", "x-fern-sdk-method-name": "get"}
```
