---
title: List firewall rules
page_id: operation-get-zones-zone-id-firewall-rules-0789ac76
path: operations/firewall-rules
description: Fetches firewall rules in a zone. You can filter the results using several optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/firewall/rules
operation_ids:
    - firewall-rules-list-firewall-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List firewall rules

`GET /zones/{zone_id}/firewall/rules`

Operation ID: `firewall-rules-list-firewall-rules`

Fetches firewall rules in a zone. You can filter the results using several optional parameters.

## Definition

```yaml
{"operationId": "firewall-rules-list-firewall-rules", "summary": "List firewall rules", "description": "Fetches firewall rules in a zone. You can filter the results using several optional parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "description", "in": "query", "schema": {"description": "A case-insensitive string to find in the description.", "type": "string", "example": "mir"}}, {"name": "action", "in": "query", "schema": {"description": "The action to search for. Must be an exact match.", "type": "string", "example": "block"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of firewall rules per page.", "type": "number", "default": 25, "maximum": 100, "minimum": 5}}, {"name": "id", "in": "query", "schema": {"description": "The unique identifier of the firewall rule.", "type": "string", "example": "372e67954025e0ba6aaa6d586b9e0b60", "maxLength": 32, "readOnly": true}}, {"name": "paused", "in": "query", "schema": {"description": "When true, indicates that the firewall rule is currently paused.", "type": "boolean", "example": false}}], "responses": {"200": {"description": "List firewall rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}}}}, "4XX": {"description": "List firewall rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
