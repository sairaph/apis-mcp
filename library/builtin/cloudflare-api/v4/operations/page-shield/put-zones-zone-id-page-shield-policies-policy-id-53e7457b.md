---
title: Update a Page Shield policy
page_id: operation-put-zones-zone-id-page-shield-policies-policy-id-da96e25e
path: operations/page-shield
description: Update a Page Shield policy by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/page_shield/policies/{policy_id}
operation_ids:
    - page-shield-update-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Page Shield policy

`PUT /zones/{zone_id}/page_shield/policies/{policy_id}`

Operation ID: `page-shield-update-policy`

Update a Page Shield policy by ID.

## Definition

```yaml
{"operationId": "page-shield-update-policy", "summary": "Update a Page Shield policy", "description": "Update a Page Shield policy by ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"action": {"$ref": "#/components/schemas/page-shield_policy-action"}, "description": {"$ref": "#/components/schemas/page-shield_policy-description"}, "enabled": {"$ref": "#/components/schemas/page-shield_policy-enabled"}, "expression": {"$ref": "#/components/schemas/page-shield_policy-expression"}, "value": {"$ref": "#/components/schemas/page-shield_policy-value"}}}}}}, "responses": {"200": {"description": "Update a Page Shield policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_get-zone-policy-response"}}}}, "4XX": {"description": "Update a Page Shield policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield", "Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
