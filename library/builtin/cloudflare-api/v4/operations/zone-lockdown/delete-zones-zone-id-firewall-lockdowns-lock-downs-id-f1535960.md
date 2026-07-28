---
title: Delete a Zone Lockdown rule
page_id: operation-delete-zones-zone-id-firewall-lockdowns-lock-downs-id-b077213b
path: operations/zone-lockdown
description: Deletes an existing Zone Lockdown rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/lockdowns/{lock_downs_id}
operation_ids:
    - zone-lockdown-delete-a-zone-lockdown-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Zone Lockdown rule

`DELETE /zones/{zone_id}/firewall/lockdowns/{lock_downs_id}`

Operation ID: `zone-lockdown-delete-a-zone-lockdown-rule`

Deletes an existing Zone Lockdown rule.

## Definition

```yaml
{"operationId": "zone-lockdown-delete-a-zone-lockdown-rule", "summary": "Delete a Zone Lockdown rule", "description": "Deletes an existing Zone Lockdown rule.", "parameters": [{"name": "lock_downs_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a Zone Lockdown rule response", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-id"}}}}}}}}, "4XX": {"description": "Delete a Zone Lockdown rule response failure", "content": {"application/json": {"schema": {"allOf": [{"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_lockdowns_components-schemas-id"}}}}, "type": "object"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Lockdown"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.lockdowns", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
