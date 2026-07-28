---
title: Update a Zone Lockdown rule
page_id: operation-put-zones-zone-id-firewall-lockdowns-lock-downs-id-36b1a955
path: operations/zone-lockdown
description: Updates an existing Zone Lockdown rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/firewall/lockdowns/{lock_downs_id}
operation_ids:
    - zone-lockdown-update-a-zone-lockdown-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Zone Lockdown rule

`PUT /zones/{zone_id}/firewall/lockdowns/{lock_downs_id}`

Operation ID: `zone-lockdown-update-a-zone-lockdown-rule`

Updates an existing Zone Lockdown rule.

## Definition

```yaml
{"operationId": "zone-lockdown-update-a-zone-lockdown-rule", "summary": "Update a Zone Lockdown rule", "description": "Updates an existing Zone Lockdown rule.", "parameters": [{"name": "lock_downs_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"configurations": {"$ref": "#/components/schemas/firewall_configurations"}, "urls": {"$ref": "#/components/schemas/firewall_urls"}}, "required": ["urls", "configurations"]}}}}, "responses": {"200": {"description": "Update a Zone Lockdown rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_zonelockdown_response_single"}}}}, "4XX": {"description": "Update a Zone Lockdown rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_zonelockdown_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Lockdown"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.lockdowns", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
