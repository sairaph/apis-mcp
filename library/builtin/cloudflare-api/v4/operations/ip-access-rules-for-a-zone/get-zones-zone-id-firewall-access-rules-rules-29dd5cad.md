---
title: List IP Access rules
page_id: operation-get-zones-zone-id-firewall-access-rules-rules-04e50249
path: operations/ip-access-rules-for-a-zone
description: Fetches IP Access rules of a zone. You can filter the results using several optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/access_rules/rules
operation_ids:
    - ip-access-rules-for-a-zone-list-ip-access-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List IP Access rules

`GET /zones/{zone_id}/firewall/access_rules/rules`

Operation ID: `ip-access-rules-for-a-zone-list-ip-access-rules`

Fetches IP Access rules of a zone. You can filter the results using several optional parameters.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-zone-list-ip-access-rules", "summary": "List IP Access rules", "description": "Fetches IP Access rules of a zone. You can filter the results using several optional parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "mode", "in": "query", "schema": {"$ref": "#/components/schemas/firewall_schemas-mode"}}, {"name": "configuration.target", "in": "query", "schema": {"description": "The target to search in existing rules.", "type": "string", "example": "ip", "enum": ["ip", "ip_range", "asn", "country"]}}, {"name": "configuration.value", "in": "query", "schema": {"description": "The target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`.\nNotes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.", "type": "string", "example": "198.51.100.4"}}, {"name": "notes", "in": "query", "schema": {"description": "The string to search for in the notes of existing IP Access rules.\nNotes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.", "type": "string", "example": "my note"}}, {"name": "match", "in": "query", "schema": {"description": "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "page", "in": "query", "schema": {"description": "Requested page within paginated list of results.", "type": "number", "example": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results requested.", "type": "number", "example": 20}}, {"name": "order", "in": "query", "schema": {"description": "The field used to sort returned rules.", "type": "string", "example": "mode", "enum": ["configuration.target", "configuration.value", "mode"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction used to sort returned rules.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List IP Access rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_collection_response"}}}}, "4XX": {"description": "List IP Access rules response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_collection_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a zone"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.zone-access-rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
