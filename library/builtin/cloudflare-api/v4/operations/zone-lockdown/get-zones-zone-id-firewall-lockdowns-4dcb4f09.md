---
title: List Zone Lockdown rules
page_id: operation-get-zones-zone-id-firewall-lockdowns-91ed3abc
path: operations/zone-lockdown
description: Fetches Zone Lockdown rules. You can filter the results using several optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/lockdowns
operation_ids:
    - zone-lockdown-list-zone-lockdown-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zone Lockdown rules

`GET /zones/{zone_id}/firewall/lockdowns`

Operation ID: `zone-lockdown-list-zone-lockdown-rules`

Fetches Zone Lockdown rules. You can filter the results using several optional parameters.

## Definition

```yaml
{"operationId": "zone-lockdown-list-zone-lockdown-rules", "summary": "List Zone Lockdown rules", "description": "Fetches Zone Lockdown rules. You can filter the results using several optional parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "description", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_schemas-description_search"}]}}, {"name": "modified_on", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_modified_on"}]}}, {"name": "ip", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_ip_search"}]}}, {"name": "priority", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_schemas-priority"}]}}, {"name": "uri_search", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_uri_search"}]}}, {"name": "ip_range_search", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/firewall_ip_range_search"}]}}, {"name": "per_page", "in": "query", "schema": {"description": "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.", "type": "number", "default": 20, "maximum": 1000, "minimum": 1}}, {"name": "created_on", "in": "query", "schema": {"description": "The timestamp of when the rule was created.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}}, {"name": "description_search", "in": "query", "schema": {"description": "A string to search for in the description of existing rules.", "type": "string", "example": "endpoints"}}, {"name": "ip_search", "in": "query", "schema": {"description": "A single IP address to search for in existing rules.", "type": "string", "example": "1.2.3.4"}}], "responses": {"200": {"description": "List Zone Lockdown rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_zonelockdown_response_collection"}}}}, "4XX": {"description": "List Zone Lockdown rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_zonelockdown_response_collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Lockdown"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.lockdowns", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
