---
title: Get Page Shield settings
page_id: operation-get-zones-zone-id-page-shield-2d99e406
path: operations/page-shield
description: Fetches the Page Shield settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield
operation_ids:
    - page-shield-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Page Shield settings

`GET /zones/{zone_id}/page_shield`

Operation ID: `page-shield-get-settings`

Fetches the Page Shield settings.

## Definition

```yaml
{"operationId": "page-shield-get-settings", "summary": "Get Page Shield settings", "description": "Fetches the Page Shield settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"200": {"description": "Get Page Shield settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/page-shield_api-get-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/page-shield_get-zone-settings-response"}}}]}}}}, "4XX": {"description": "Get Page Shield settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
