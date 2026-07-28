---
title: Get a WAF rule
page_id: operation-get-zones-zone-id-firewall-waf-packages-package-id-rules-rule-id-4f38e54a
path: operations/waf-rules
description: |-
    Fetches the details of a WAF rule in a WAF package.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}
operation_ids:
    - waf-rules-get-a-waf-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a WAF rule

`GET /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}`

Operation ID: `waf-rules-get-a-waf-rule`

Fetches the details of a WAF rule in a WAF package.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-rules-get-a-waf-rule", "summary": "Get a WAF rule", "description": "Fetches the details of a WAF rule in a WAF package.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_schemas-identifier"}}], "responses": {"200": {"description": "Get a WAF rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-managed-rules_rule_response_single"}}}}, "4XX": {"description": "Get a WAF rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_response_single"}, {"$ref": "#/components/schemas/waf-managed-rules_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF rules"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
