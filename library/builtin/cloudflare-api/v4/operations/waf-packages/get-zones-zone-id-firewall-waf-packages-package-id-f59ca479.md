---
title: Get a WAF package
page_id: operation-get-zones-zone-id-firewall-waf-packages-package-id-24f8d5cd
path: operations/waf-packages
description: |-
    Fetches the details of a WAF package.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}
operation_ids:
    - waf-packages-get-a-waf-package
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a WAF package

`GET /zones/{zone_id}/firewall/waf/packages/{package_id}`

Operation ID: `waf-packages-get-a-waf-package`

Fetches the details of a WAF package.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-packages-get-a-waf-package", "summary": "Get a WAF package", "description": "Fetches the details of a WAF package.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_package_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "responses": {"200": {"description": "Get a WAF package response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_package_response_single"}}}}, "4XX": {"description": "Get a WAF package response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_package_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF packages"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.waf.packages", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
