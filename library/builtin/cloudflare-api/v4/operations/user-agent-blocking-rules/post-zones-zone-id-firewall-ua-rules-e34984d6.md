---
title: Create a User Agent Blocking rule
page_id: operation-post-zones-zone-id-firewall-ua-rules-d3f9c2cf
path: operations/user-agent-blocking-rules
description: Creates a new User Agent Blocking rule in a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/firewall/ua_rules
operation_ids:
    - user-agent-blocking-rules-create-a-user-agent-blocking-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a User Agent Blocking rule

`POST /zones/{zone_id}/firewall/ua_rules`

Operation ID: `user-agent-blocking-rules-create-a-user-agent-blocking-rule`

Creates a new User Agent Blocking rule in a zone.

## Definition

```yaml
{"operationId": "user-agent-blocking-rules-create-a-user-agent-blocking-rule", "summary": "Create a User Agent Blocking rule", "description": "Creates a new User Agent Blocking rule in a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/firewall_ua_configuration"}, "description": {"$ref": "#/components/schemas/firewall_description"}, "mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "paused": {"$ref": "#/components/schemas/firewall_schemas-paused"}}, "required": ["mode", "configuration"]}}}}, "responses": {"200": {"description": "Create a User Agent Blocking rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}}}}, "4XX": {"description": "Create a User Agent Blocking rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Agent Blocking rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.ua-rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
