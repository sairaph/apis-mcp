---
title: Create a WAF override
page_id: operation-post-zones-zone-id-firewall-waf-overrides-1601a58b
path: operations/waf-overrides
description: |-
    **This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

    Previously created a URI-based WAF override for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/firewall/waf/overrides
operation_ids:
    - waf-overrides-create-a-waf-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a WAF override

`POST /zones/{zone_id}/firewall/waf/overrides`

Operation ID: `waf-overrides-create-a-waf-override`

**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

Previously created a URI-based WAF override for a zone.

## Definition

```yaml
{"operationId": "waf-overrides-create-a-waf-override", "summary": "Create a WAF override", "description": "**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**\n\nPreviously created a URI-based WAF override for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"urls": {"$ref": "#/components/schemas/firewall_urls"}}, "required": ["urls"]}}}}, "responses": {"200": {"description": "Create a WAF override response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_override_response_single"}}}}, "4XX": {"description": "Create a WAF override response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_override_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF overrides"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.overrides", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
