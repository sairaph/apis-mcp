---
title: Update a WAF rule group
page_id: operation-patch-zones-zone-id-firewall-waf-packages-package-id-groups-group-id-abc61acf
path: operations/waf-rule-groups
description: |-
    Updates a WAF rule group. You can update the state (`mode` parameter) of a rule group.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}
operation_ids:
    - waf-rule-groups-update-a-waf-rule-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a WAF rule group

`PATCH /zones/{zone_id}/firewall/waf/packages/{package_id}/groups/{group_id}`

Operation ID: `waf-rule-groups-update-a-waf-rule-group`

Updates a WAF rule group. You can update the state (`mode` parameter) of a rule group.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-rule-groups-update-a-waf-rule-group", "summary": "Update a WAF rule group", "description": "Updates a WAF rule group. You can update the state (`mode` parameter) of a rule group.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"mode": {"$ref": "#/components/schemas/waf-managed-rules_mode"}}}}}}, "responses": {"200": {"description": "Update a WAF rule group response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-managed-rules_rule_group_response_single"}}}}, "4XX": {"description": "Update a WAF rule group response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_group_response_single"}, {"$ref": "#/components/schemas/waf-managed-rules_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF rule groups"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
