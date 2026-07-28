---
title: List WAF overrides
page_id: operation-get-zones-zone-id-firewall-waf-overrides-0e19bdf4
path: operations/waf-overrides
description: |-
    **This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

    Previously fetched the URI-based WAF overrides in a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/overrides
operation_ids:
    - waf-overrides-list-waf-overrides
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WAF overrides

`GET /zones/{zone_id}/firewall/waf/overrides`

Operation ID: `waf-overrides-list-waf-overrides`

**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

Previously fetched the URI-based WAF overrides in a zone.

## Definition

```yaml
{"operationId": "waf-overrides-list-waf-overrides", "summary": "List WAF overrides", "description": "**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**\n\nPreviously fetched the URI-based WAF overrides in a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "The page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "The number of WAF overrides per page.", "type": "number", "default": 50, "maximum": 100, "minimum": 5}}], "responses": {"200": {"description": "List WAF overrides response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_override_response_collection"}}}}, "4XX": {"description": "List WAF overrides response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_override_response_collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF overrides"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.overrides", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
