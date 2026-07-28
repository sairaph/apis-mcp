---
title: Create a Zone Lockdown rule
page_id: operation-post-zones-zone-id-firewall-lockdowns-b96b07db
path: operations/zone-lockdown
description: Creates a new Zone Lockdown rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/firewall/lockdowns
operation_ids:
    - zone-lockdown-create-a-zone-lockdown-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Zone Lockdown rule

`POST /zones/{zone_id}/firewall/lockdowns`

Operation ID: `zone-lockdown-create-a-zone-lockdown-rule`

Creates a new Zone Lockdown rule.

## Definition

```yaml
{"operationId": "zone-lockdown-create-a-zone-lockdown-rule", "summary": "Create a Zone Lockdown rule", "description": "Creates a new Zone Lockdown rule.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"configurations": {"$ref": "#/components/schemas/firewall_configurations"}, "description": {"$ref": "#/components/schemas/firewall_description"}, "paused": {"$ref": "#/components/schemas/firewall_schemas-paused"}, "priority": {"$ref": "#/components/schemas/firewall_schemas-priority"}, "urls": {"$ref": "#/components/schemas/firewall_urls"}}, "required": ["urls", "configurations"]}}}}, "responses": {"200": {"description": "Create a Zone Lockdown rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_zonelockdown_response_single"}}}}, "4XX": {"description": "Create a Zone Lockdown rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_zonelockdown_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Lockdown"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.lockdowns", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
