---
title: Create firewall rules
page_id: operation-post-zones-zone-id-firewall-rules-690df267
path: operations/firewall-rules
description: Create one or more firewall rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/firewall/rules
operation_ids:
    - firewall-rules-create-firewall-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create firewall rules

`POST /zones/{zone_id}/firewall/rules`

Operation ID: `firewall-rules-create-firewall-rules`

Create one or more firewall rules.

## Definition

```yaml
{"operationId": "firewall-rules-create-firewall-rules", "summary": "Create firewall rules", "description": "Create one or more firewall rules.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"$ref": "#/components/schemas/firewall_action"}, "filter": {"$ref": "#/components/schemas/firewall_filter"}}, "required": ["filter", "action"]}}}}, "responses": {"200": {"description": "Create firewall rules response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}}}}, "4XX": {"description": "Create firewall rules response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-rules-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Firewall rules"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
