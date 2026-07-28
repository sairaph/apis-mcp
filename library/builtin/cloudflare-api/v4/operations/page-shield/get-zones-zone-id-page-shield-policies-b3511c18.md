---
title: List Page Shield policies
page_id: operation-get-zones-zone-id-page-shield-policies-5756a203
path: operations/page-shield
description: Lists all Page Shield policies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/policies
operation_ids:
    - page-shield-list-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Page Shield policies

`GET /zones/{zone_id}/page_shield/policies`

Operation ID: `page-shield-list-policies`

Lists all Page Shield policies.

## Definition

```yaml
{"operationId": "page-shield-list-policies", "summary": "List Page Shield policies", "description": "Lists all Page Shield policies.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"200": {"description": "List Page Shield policies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_list-zone-policies-response"}}}}, "4XX": {"description": "List Page Shield policies response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
