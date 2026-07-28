---
title: Update a WAF rule
page_id: operation-patch-zones-zone-id-firewall-waf-packages-package-id-rules-rule-id-0c35fe16
path: operations/waf-rules
description: |-
    Updates a WAF rule. You can only update the mode/action of the rule.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}
operation_ids:
    - waf-rules-update-a-waf-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a WAF rule

`PATCH /zones/{zone_id}/firewall/waf/packages/{package_id}/rules/{rule_id}`

Operation ID: `waf-rules-update-a-waf-rule`

Updates a WAF rule. You can only update the mode/action of the rule.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-rules-update-a-waf-rule", "summary": "Update a WAF rule", "description": "Updates a WAF rule. You can only update the mode/action of the rule.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"mode": {"description": "Defines the mode/action of the rule when triggered. You must use a value from the `allowed_modes` array of the current rule.", "type": "string", "example": "on", "enum": ["default", "disable", "simulate", "block", "challenge", "on", "off"]}}}}}}, "responses": {"200": {"description": "Update a WAF rule response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_response_single"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/waf-managed-rules_anomaly_rule"}, {"$ref": "#/components/schemas/waf-managed-rules_traditional_deny_rule"}, {"$ref": "#/components/schemas/waf-managed-rules_traditional_allow_rule"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Update a WAF rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_response_single"}, {"properties": {"result": {"oneOf": [{"$ref": "#/components/schemas/waf-managed-rules_anomaly_rule"}, {"$ref": "#/components/schemas/waf-managed-rules_traditional_deny_rule"}, {"$ref": "#/components/schemas/waf-managed-rules_traditional_allow_rule"}]}}, "type": "object"}, {"$ref": "#/components/schemas/waf-managed-rules_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
