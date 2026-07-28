---
title: List WAF packages
page_id: operation-get-zones-zone-id-firewall-waf-packages-1c4097a4
path: operations/waf-packages
description: |-
    Fetches WAF packages for a zone.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages
operation_ids:
    - waf-packages-list-waf-packages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WAF packages

`GET /zones/{zone_id}/firewall/waf/packages`

Operation ID: `waf-packages-list-waf-packages`

Fetches WAF packages for a zone.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-packages-list-waf-packages", "summary": "List WAF packages", "description": "Fetches WAF packages for a zone.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "The page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "The number of packages per page.", "type": "number", "default": 50, "maximum": 100, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "The field used to sort returned packages.", "type": "string", "example": "name", "enum": ["name"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction used to sort returned packages.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "name", "in": "query", "schema": {"description": "The name of the WAF package.", "type": "string", "example": "USER", "readOnly": true}}], "responses": {"200": {"description": "List WAF packages response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_package_response_collection"}}}}, "4XX": {"description": "List WAF packages response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_package_response_collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF packages"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.packages", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
