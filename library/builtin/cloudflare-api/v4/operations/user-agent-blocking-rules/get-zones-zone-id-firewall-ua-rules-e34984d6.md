---
title: List User Agent Blocking rules
page_id: operation-get-zones-zone-id-firewall-ua-rules-dbb7d0ce
path: operations/user-agent-blocking-rules
description: Fetches User Agent Blocking rules in a zone. You can filter the results using several optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/ua_rules
operation_ids:
    - user-agent-blocking-rules-list-user-agent-blocking-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List User Agent Blocking rules

`GET /zones/{zone_id}/firewall/ua_rules`

Operation ID: `user-agent-blocking-rules-list-user-agent-blocking-rules`

Fetches User Agent Blocking rules in a zone. You can filter the results using several optional parameters.

## Definition

```yaml
{"operationId": "user-agent-blocking-rules-list-user-agent-blocking-rules", "summary": "List User Agent Blocking rules", "description": "Fetches User Agent Blocking rules in a zone. You can filter the results using several optional parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "description", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_description_search"}]}}, {"name": "per_page", "in": "query", "schema": {"description": "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.", "type": "number", "default": 20, "maximum": 1000, "minimum": 1}}, {"name": "user_agent", "in": "query", "schema": {"description": "A string to search for in the user agent values of existing rules.", "type": "string", "example": "Safari"}}, {"name": "paused", "in": "query", "schema": {"description": "When true, indicates that the rule is currently paused.", "type": "boolean", "example": false}}], "responses": {"200": {"description": "List User Agent Blocking rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_firewalluablock_response_collection"}}}}, "4XX": {"description": "List User Agent Blocking rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_firewalluablock_response_collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Agent Blocking rules"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.ua-rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
