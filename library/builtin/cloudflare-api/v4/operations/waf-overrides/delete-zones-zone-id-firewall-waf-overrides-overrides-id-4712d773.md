---
title: Delete a WAF override
page_id: operation-delete-zones-zone-id-firewall-waf-overrides-overrides-id-dce07336
path: operations/waf-overrides
description: |-
    **This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

    Previously deleted an existing URI-based WAF override.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/waf/overrides/{overrides_id}
operation_ids:
    - waf-overrides-delete-a-waf-override
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a WAF override

`DELETE /zones/{zone_id}/firewall/waf/overrides/{overrides_id}`

Operation ID: `waf-overrides-delete-a-waf-override`

**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**

Previously deleted an existing URI-based WAF override.

## Definition

```yaml
{"operationId": "waf-overrides-delete-a-waf-override", "summary": "Delete a WAF override", "description": "**This endpoint has been deprecated and returns 410 Gone. Please use the [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**\n\nPreviously deleted an existing URI-based WAF override.", "parameters": [{"name": "overrides_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_overrides-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a WAF override response", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_overrides-id"}}}}}}}}, "4XX": {"description": "Delete a WAF override response failure", "content": {"application/json": {"schema": {"allOf": [{"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_overrides-id"}}}}, "type": "object"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF overrides"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.overrides", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
