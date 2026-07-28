---
title: List WAF rule groups
page_id: operation-get-zones-zone-id-firewall-waf-packages-package-id-groups-0d6028f9
path: operations/waf-rule-groups
description: |-
    Fetches the WAF rule groups in a WAF package.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}/groups
operation_ids:
    - waf-rule-groups-list-waf-rule-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WAF rule groups

`GET /zones/{zone_id}/firewall/waf/packages/{package_id}/groups`

Operation ID: `waf-rule-groups-list-waf-rule-groups`

Fetches the WAF rule groups in a WAF package.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-rule-groups-list-waf-rule-groups", "summary": "List WAF rule groups", "description": "Fetches the WAF rule groups in a WAF package.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_schemas-identifier"}}, {"name": "mode", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_mode"}]}}, {"name": "page", "in": "query", "schema": {"description": "Defines the page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Defines the number of rule groups per page.", "type": "number", "default": 50, "maximum": 100, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Defines the field used to sort returned rule groups.", "type": "string", "example": "mode", "enum": ["mode", "rules_count"]}}, {"name": "direction", "in": "query", "schema": {"description": "Defines the direction used to sort returned rule groups.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "Defines the condition for search requirements. When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "name", "in": "query", "schema": {"description": "Defines the name of the rule group.", "type": "string", "example": "Project Honey Pot", "readOnly": true}}, {"name": "rules_count", "in": "query", "schema": {"description": "Defines the number of rules in the current rule group.", "type": "number", "example": 10, "default": 0, "readOnly": true}}], "responses": {"200": {"description": "Defines the list WAF rule groups response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-managed-rules_rule_group_response_collection"}}}}, "4XX": {"description": "Defines the list WAF rule groups response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_group_response_collection"}, {"$ref": "#/components/schemas/waf-managed-rules_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF rule groups"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
