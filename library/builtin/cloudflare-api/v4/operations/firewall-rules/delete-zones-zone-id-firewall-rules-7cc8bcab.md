---
title: Delete firewall rules
page_id: operation-delete-zones-zone-id-firewall-rules-fd9a5bec
path: operations/firewall-rules
description: Deletes existing firewall rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/rules
operation_ids:
    - firewall-rules-delete-firewall-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete firewall rules

`DELETE /zones/{zone_id}/firewall/rules`

Operation ID: `firewall-rules-delete-firewall-rules`

Deletes existing firewall rules.

## Definition

```yaml
{"operationId": "firewall-rules-delete-firewall-rules", "summary": "Delete firewall rules", "description": "Deletes existing firewall rules.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/firewall_firewall-rules_components-schemas-id"}}, "required": ["id"]}}}}, "responses": {"200": {"description": "Delete firewall rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-response-collection-delete"}}}}, "4XX": {"description": "Delete firewall rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-response-collection-delete"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
