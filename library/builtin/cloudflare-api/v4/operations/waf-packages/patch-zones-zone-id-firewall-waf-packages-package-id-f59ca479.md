---
title: Update a WAF package
page_id: operation-patch-zones-zone-id-firewall-waf-packages-package-id-7bd89438
path: operations/waf-packages
description: |-
    Updates a WAF package. You can update the sensitivity and the action of an anomaly detection WAF package.

    **Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/firewall/waf/packages/{package_id}
operation_ids:
    - waf-packages-update-a-waf-package
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a WAF package

`PATCH /zones/{zone_id}/firewall/waf/packages/{package_id}`

Operation ID: `waf-packages-update-a-waf-package`

Updates a WAF package. You can update the sensitivity and the action of an anomaly detection WAF package.

**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).

## Definition

```yaml
{"operationId": "waf-packages-update-a-waf-package", "summary": "Update a WAF package", "description": "Updates a WAF package. You can update the sensitivity and the action of an anomaly detection WAF package.\n\n**Note:** Applies only to the [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).", "parameters": [{"name": "package_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_package_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action_mode": {"$ref": "#/components/schemas/firewall_action_mode"}, "sensitivity": {"$ref": "#/components/schemas/firewall_sensitivity"}}}}}}, "responses": {"200": {"description": "Update a WAF package response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_package_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/firewall_anomaly_package"}}, "type": "object"}]}}}}, "4XX": {"description": "Update a WAF package response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/firewall_package_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/firewall_anomaly_package"}}, "type": "object"}]}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["WAF packages"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "firewall.waf.packages", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}}
```
