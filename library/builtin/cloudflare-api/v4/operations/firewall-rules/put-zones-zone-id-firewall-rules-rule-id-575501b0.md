---
title: Update a firewall rule
page_id: operation-put-zones-zone-id-firewall-rules-rule-id-782d3aa8
path: operations/firewall-rules
description: Updates an existing firewall rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/firewall/rules/{rule_id}
operation_ids:
    - firewall-rules-update-a-firewall-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a firewall rule

`PUT /zones/{zone_id}/firewall/rules/{rule_id}`

Operation ID: `firewall-rules-update-a-firewall-rule`

Updates an existing firewall rule.

## Definition

```yaml
{"operationId": "firewall-rules-update-a-firewall-rule", "summary": "Update a firewall rule", "description": "Updates an existing firewall rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_firewall-rules_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"$ref": "#/components/schemas/firewall_action"}, "filter": {"$ref": "#/components/schemas/firewall_filter"}, "id": {"$ref": "#/components/schemas/firewall_components-schemas-identifier"}}, "required": ["id", "filter", "action"]}}}}, "responses": {"200": {"description": "Update a firewall rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-single-response"}}}}, "4XX": {"description": "Update a firewall rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-single-response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
