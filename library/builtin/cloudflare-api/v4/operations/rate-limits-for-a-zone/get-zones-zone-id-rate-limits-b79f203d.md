---
title: List rate limits
page_id: operation-get-zones-zone-id-rate-limits-36ae9899
path: operations/rate-limits-for-a-zone
description: '**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/rate_limits
operation_ids:
    - rate-limits-for-a-zone-list-rate-limits
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List rate limits

`GET /zones/{zone_id}/rate_limits`

Operation ID: `rate-limits-for-a-zone-list-rate-limits`

**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.

## Definition

```yaml
{"operationId": "rate-limits-for-a-zone-list-rate-limits", "summary": "List rate limits", "description": "**Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API instead.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Defines the page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Defines the maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.", "type": "number", "default": 20, "maximum": 1000, "minimum": 1}}], "responses": {"200": {"description": "List rate limits response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_ratelimit_response_collection"}}}}, "4XX": {"description": "List rate limits response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_ratelimit_response_collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Rate limits for a zone"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rate-limits", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
