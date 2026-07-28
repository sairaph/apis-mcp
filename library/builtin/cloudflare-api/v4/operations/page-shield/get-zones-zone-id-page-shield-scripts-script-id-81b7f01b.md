---
title: Get a Page Shield script
page_id: operation-get-zones-zone-id-page-shield-scripts-script-id-c1737305
path: operations/page-shield
description: Fetches a script detected by Page Shield by script ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/scripts/{script_id}
operation_ids:
    - page-shield-get-script
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Page Shield script

`GET /zones/{zone_id}/page_shield/scripts/{script_id}`

Operation ID: `page-shield-get-script`

Fetches a script detected by Page Shield by script ID.

## Definition

```yaml
{"operationId": "page-shield-get-script", "summary": "Get a Page Shield script", "description": "Fetches a script detected by Page Shield by script ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "script_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"200": {"description": "Get a Page Shield script response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_get-zone-script-response"}}}}, "4XX": {"description": "Get a Page Shield script response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/page-shield_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
