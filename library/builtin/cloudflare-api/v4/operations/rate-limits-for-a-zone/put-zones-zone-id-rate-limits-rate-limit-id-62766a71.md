---
title: Update a rate limit
page_id: operation-put-zones-zone-id-rate-limits-rate-limit-id-39b5956e
path: operations/rate-limits-for-a-zone
description: '**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/rate_limits/{rate_limit_id}
operation_ids:
    - rate-limits-for-a-zone-update-a-rate-limit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a rate limit

`PUT /zones/{zone_id}/rate_limits/{rate_limit_id}`

Operation ID: `rate-limits-for-a-zone-update-a-rate-limit`

**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.

## Definition

```yaml
{"operationId": "rate-limits-for-a-zone-update-a-rate-limit", "summary": "Update a rate limit", "description": "**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.", "parameters": [{"name": "rate_limit_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rate_limit_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"$ref": "#/components/schemas/firewall_action"}, "match": {"$ref": "#/components/schemas/firewall_match"}, "period": {"$ref": "#/components/schemas/firewall_period"}, "threshold": {"$ref": "#/components/schemas/firewall_threshold"}}, "required": ["match", "threshold", "period", "action"]}}}}, "responses": {"200": {"description": "Update a rate limit response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_ratelimit_response_single"}}}}, "4XX": {"description": "Update a rate limit response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_ratelimit_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Rate limits for a zone"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rate-limits", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
