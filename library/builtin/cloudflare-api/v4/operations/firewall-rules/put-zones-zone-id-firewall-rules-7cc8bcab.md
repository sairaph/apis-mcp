---
title: Update firewall rules
page_id: operation-put-zones-zone-id-firewall-rules-9d87b74e
path: operations/firewall-rules
description: Updates one or more existing firewall rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/firewall/rules
operation_ids:
    - firewall-rules-update-firewall-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update firewall rules

`PUT /zones/{zone_id}/firewall/rules`

Operation ID: `firewall-rules-update-firewall-rules`

Updates one or more existing firewall rules.

## Definition

```yaml
{"operationId": "firewall-rules-update-firewall-rules", "summary": "Update firewall rules", "description": "Updates one or more existing firewall rules.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"required": ["id"]}}}}, "responses": {"200": {"description": "Update firewall rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}}}}, "4XX": {"description": "Update firewall rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "bulk-update", "x-forge-hidden": true}
```
