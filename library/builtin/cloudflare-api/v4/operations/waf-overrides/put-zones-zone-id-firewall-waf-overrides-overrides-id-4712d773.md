---
title: Update WAF override
page_id: operation-put-zones-zone-id-firewall-waf-overrides-overrides-id-0aa102e2
path: operations/waf-overrides
description: |-
    **This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

    Previously updated an existing URI-based WAF override.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/firewall/waf/overrides/{overrides_id}
operation_ids:
    - waf-overrides-update-waf-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update WAF override

`PUT /zones/{zone_id}/firewall/waf/overrides/{overrides_id}`

Operation ID: `waf-overrides-update-waf-override`

**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

Previously updated an existing URI-based WAF override.

## Definition

```yaml
{"operationId": "waf-overrides-update-waf-override", "summary": "Update WAF override", "description": "**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**\n\nPreviously updated an existing URI-based WAF override.", "parameters": [{"name": "overrides_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_overrides-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_identifier"}, "rewrite_action": {"$ref": "#/components/schemas/firewall_rewrite_action"}, "rules": {"$ref": "#/components/schemas/firewall_rules"}, "urls": {"$ref": "#/components/schemas/firewall_urls"}}, "required": ["id", "urls", "rules", "rewrite_action"]}}}}, "responses": {"200": {"description": "Update WAF override response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_override_response_single"}}}}, "4XX": {"description": "Update WAF override response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_override_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF overrides"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.overrides", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
