---
title: Available Plan Details
page_id: operation-get-zones-zone-id-available-plans-plan-identifier-fbc29140
path: operations/zone-rate-plan
description: Details of the available plan that the zone can subscribe to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/available_plans/{plan_identifier}
operation_ids:
    - zone-rate-plan-available-plan-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Available Plan Details

`GET /zones/{zone_id}/available_plans/{plan_identifier}`

Operation ID: `zone-rate-plan-available-plan-details`

Details of the available plan that the zone can subscribe to.

## Definition

```yaml
{"operationId": "zone-rate-plan-available-plan-details", "summary": "Available Plan Details", "description": "Details of the available plan that the zone can subscribe to.", "parameters": [{"name": "plan_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "Available Plan Details response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/bill-subs-api_available-rate-plan"}}}]}}}}, "4XX": {"description": "Available Plan Details response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/bill-subs-api_available-rate-plan"}}}]}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Rate Plan"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.plans", "x-fern-sdk-method-name": "get"}
```
