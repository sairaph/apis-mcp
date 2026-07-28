---
title: List WAF rules
page_id: operation-get-zones-zone-id-firewall-waf-packages-package-id-rules-fed285f6
path: operations/waf-rules
description: |-
    Fetches WAF rules in a WAF package.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}/rules
operation_ids:
    - waf-rules-list-waf-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WAF rules

`GET /zones/{zone_id}/firewall/waf/packages/{package_id}/rules`

Operation ID: `waf-rules-list-waf-rules`

Fetches WAF rules in a WAF package.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-rules-list-waf-rules", "summary": "List WAF rules", "description": "Fetches WAF rules in a WAF package.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-managed-rules_schemas-identifier"}}, {"name": "mode", "in": "query", "schema": {"description": "Defines the action/mode a rule has been overridden to perform.", "type": "string", "example": "CHL", "enum": ["DIS", "CHL", "BLK", "SIM"]}}, {"name": "group_id", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_components-schemas-identifier"}]}}, {"name": "page", "in": "query", "schema": {"description": "Defines the page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Defines the number of rules per page.", "type": "number", "default": 50, "maximum": 100, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Defines the field used to sort returned rules.", "type": "string", "example": "status", "enum": ["priority", "group_id", "description"]}}, {"name": "direction", "in": "query", "schema": {"description": "Defines the direction used to sort returned rules.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "Defines the search requirements. When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "description", "in": "query", "schema": {"description": "Defines the public description of the WAF rule.", "type": "string", "example": "SQL injection prevention for SELECT statements", "readOnly": true}}, {"name": "priority", "in": "query", "schema": {"description": "Defines the order in which the individual WAF rule is executed within its rule group.", "type": "string", "readOnly": true}}], "responses": {"200": {"description": "List WAF rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-managed-rules_rule_response_collection"}}}}, "4XX": {"description": "List WAF rules response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-managed-rules_rule_response_collection"}, {"$ref": "#/components/schemas/waf-managed-rules_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF rules"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
