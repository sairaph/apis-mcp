---
title: Delete a firewall rule
page_id: operation-delete-zones-zone-id-firewall-rules-rule-id-94b18e9a
path: operations/firewall-rules
description: Deletes an existing firewall rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/rules/{rule_id}
operation_ids:
    - firewall-rules-delete-a-firewall-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a firewall rule

`DELETE /zones/{zone_id}/firewall/rules/{rule_id}`

Operation ID: `firewall-rules-delete-a-firewall-rule`

Deletes an existing firewall rule.

## Definition

```yaml
{"operationId": "firewall-rules-delete-a-firewall-rule", "summary": "Delete a firewall rule", "description": "Deletes an existing firewall rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_firewall-rules_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"delete_filter_if_unused": {"$ref": "#/components/schemas/firewall_delete_filter_if_unused"}}}}}}, "responses": {"200": {"description": "Delete a firewall rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-single-response-delete"}}}}, "4XX": {"description": "Delete a firewall rule response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-single-response-delete"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
