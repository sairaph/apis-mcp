---
title: Create a Page Shield policy
page_id: operation-post-zones-zone-id-page-shield-policies-1c317f79
path: operations/page-shield
description: Create a Page Shield policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/page_shield/policies
operation_ids:
    - page-shield-create-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Page Shield policy

`POST /zones/{zone_id}/page_shield/policies`

Operation ID: `page-shield-create-policy`

Create a Page Shield policy.

## Definition

```yaml
{"operationId": "page-shield-create-policy", "summary": "Create a Page Shield policy", "description": "Create a Page Shield policy.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_policy"}}}}, "responses": {"200": {"description": "Create a Page Shield policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_get-zone-policy-response"}}}}, "4XX": {"description": "Create a Page Shield policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield", "Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
