---
title: Get a WAF override
page_id: operation-get-zones-zone-id-firewall-waf-overrides-overrides-id-6090d573
path: operations/waf-overrides
description: |-
    **This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

    Previously fetched the details of a URI-based WAF override.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/overrides/{overrides_id}
operation_ids:
    - waf-overrides-get-a-waf-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a WAF override

`GET /zones/{zone_id}/firewall/waf/overrides/{overrides_id}`

Operation ID: `waf-overrides-get-a-waf-override`

**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

Previously fetched the details of a URI-based WAF override.

## Definition

```yaml
{"operationId": "waf-overrides-get-a-waf-override", "summary": "Get a WAF override", "description": "**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**\n\nPreviously fetched the details of a URI-based WAF override.", "parameters": [{"name": "overrides_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_overrides-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "responses": {"200": {"description": "Get a WAF override response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_override_response_single"}}}}, "4XX": {"description": "Get a WAF override response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_override_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF overrides"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.overrides", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
