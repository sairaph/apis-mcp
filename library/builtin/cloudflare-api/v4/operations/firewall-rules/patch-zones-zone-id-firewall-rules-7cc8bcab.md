---
title: Update priority of firewall rules
page_id: operation-patch-zones-zone-id-firewall-rules-60a8b471
path: operations/firewall-rules
description: Updates the priority of existing firewall rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/firewall/rules
operation_ids:
    - firewall-rules-update-priority-of-firewall-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update priority of firewall rules

`PATCH /zones/{zone_id}/firewall/rules`

Operation ID: `firewall-rules-update-priority-of-firewall-rules`

Updates the priority of existing firewall rules.

## Definition

```yaml
{"operationId": "firewall-rules-update-priority-of-firewall-rules", "summary": "Update priority of firewall rules", "description": "Updates the priority of existing firewall rules.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"required": ["id"]}}}}, "responses": {"200": {"description": "Update priority of firewall rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}}}}, "4XX": {"description": "Update priority of firewall rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "bulk-edit", "x-forge-hidden": true}
```
