---
title: Delete a User Agent Blocking rule
page_id: operation-delete-zones-zone-id-firewall-ua-rules-ua-rule-id-5f4ba041
path: operations/user-agent-blocking-rules
description: Deletes an existing User Agent Blocking rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/ua_rules/{ua_rule_id}
operation_ids:
    - user-agent-blocking-rules-delete-a-user-agent-blocking-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a User Agent Blocking rule

`DELETE /zones/{zone_id}/firewall/ua_rules/{ua_rule_id}`

Operation ID: `user-agent-blocking-rules-delete-a-user-agent-blocking-rule`

Deletes an existing User Agent Blocking rule.

## Definition

```yaml
{"operationId": "user-agent-blocking-rules-delete-a-user-agent-blocking-rule", "summary": "Delete a User Agent Blocking rule", "description": "Deletes an existing User Agent Blocking rule.", "parameters": [{"name": "ua_rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_components-ua-rule-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a User Agent Blocking rule response", "content": {"application/json": {"schema": {"allOf": [{"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_components-ua-rule-id"}}}}, "type": "object"}, {"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}]}}}}, "4XX": {"description": "Delete a User Agent Blocking rule response failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_components-ua-rule-id"}}}}, "type": "object"}]}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Agent Blocking rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.ua-rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
