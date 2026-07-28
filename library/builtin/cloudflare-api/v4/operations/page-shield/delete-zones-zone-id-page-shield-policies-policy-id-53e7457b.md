---
title: Delete a Page Shield policy
page_id: operation-delete-zones-zone-id-page-shield-policies-policy-id-cc977d4d
path: operations/page-shield
description: Delete a Page Shield policy by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/page_shield/policies/{policy_id}
operation_ids:
    - page-shield-delete-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Page Shield policy

`DELETE /zones/{zone_id}/page_shield/policies/{policy_id}`

Operation ID: `page-shield-delete-policy`

Delete a Page Shield policy by ID.

## Definition

```yaml
{"operationId": "page-shield-delete-policy", "summary": "Delete a Page Shield policy", "description": "Delete a Page Shield policy by ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"204": {"description": "Delete a Page Shield policy response"}, "4XX": {"description": "Delete a Page Shield policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield", "Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
