---
title: Get a Page Shield connection
page_id: operation-get-zones-zone-id-page-shield-connections-connection-id-5c6bda4f
path: operations/page-shield
description: Fetches a connection detected by Page Shield by connection ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/connections/{connection_id}
operation_ids:
    - page-shield-get-connection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Page Shield connection

`GET /zones/{zone_id}/page_shield/connections/{connection_id}`

Operation ID: `page-shield-get-connection`

Fetches a connection detected by Page Shield by connection ID.

## Definition

```yaml
{"operationId": "page-shield-get-connection", "summary": "Get a Page Shield connection", "description": "Fetches a connection detected by Page Shield by connection ID.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "connection_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}], "responses": {"200": {"description": "Get a Page Shield connection response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_get-zone-connection-response"}}}}, "4XX": {"description": "Get a Page Shield connection response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
