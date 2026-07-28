---
title: Get RUM status for a zone
page_id: operation-get-zones-zone-id-settings-rum-7b1f16f3
path: operations/web-analytics
description: Retrieves RUM status for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/settings/rum
operation_ids:
    - web-analytics-get-rum-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get RUM status for a zone

`GET /zones/{zone_id}/settings/rum`

Operation ID: `web-analytics-get-rum-status`

Retrieves RUM status for a zone.

## Definition

```yaml
{"operationId": "web-analytics-get-rum-status", "summary": "Get RUM status for a zone", "description": "Retrieves RUM status for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "responses": {"200": {"description": "Rum Status.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rum-site-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.rum", "x-fern-sdk-method-name": "get"}
```
