---
title: Get a Page Shield cookie
page_id: operation-get-zones-zone-id-page-shield-cookies-cookie-id-d28fe7d1
path: operations/page-shield
description: Fetches a cookie collected by Page Shield by cookie ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/cookies/{cookie_id}
operation_ids:
    - page-shield-get-cookie
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Page Shield cookie

`GET /zones/{zone_id}/page_shield/cookies/{cookie_id}`

Operation ID: `page-shield-get-cookie`

Fetches a cookie collected by Page Shield by cookie ID.

## Definition

```yaml
{"operationId": "page-shield-get-cookie", "summary": "Get a Page Shield cookie", "description": "Fetches a cookie collected by Page Shield by cookie ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "cookie_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"200": {"description": "Get a Page Shield cookie response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_get-zone-cookie-response"}}}}, "4XX": {"description": "Get a Page Shield cookie response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
