---
title: Get a User Agent Blocking rule
page_id: operation-get-zones-zone-id-firewall-ua-rules-ua-rule-id-d0df58ab
path: operations/user-agent-blocking-rules
description: Fetches the details of a User Agent Blocking rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/ua_rules/{ua_rule_id}
operation_ids:
    - user-agent-blocking-rules-get-a-user-agent-blocking-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a User Agent Blocking rule

`GET /zones/{zone_id}/firewall/ua_rules/{ua_rule_id}`

Operation ID: `user-agent-blocking-rules-get-a-user-agent-blocking-rule`

Fetches the details of a User Agent Blocking rule.

## Definition

```yaml
{"operationId": "user-agent-blocking-rules-get-a-user-agent-blocking-rule", "summary": "Get a User Agent Blocking rule", "description": "Fetches the details of a User Agent Blocking rule.", "parameters": [{"name": "ua_rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_components-ua-rule-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "responses": {"200": {"description": "Get a User Agent Blocking rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}}}}, "4XX": {"description": "Get a User Agent Blocking rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_firewalluablock_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Agent Blocking rules"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.ua-rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
