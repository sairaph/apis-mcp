---
title: Toggle RUM on/off for a zone
page_id: operation-patch-zones-zone-id-settings-rum-3af6eb9f
path: operations/web-analytics
description: Toggles RUM on/off for an existing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/rum
operation_ids:
    - web-analytics-toggle-rum
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Toggle RUM on/off for a zone

`PATCH /zones/{zone_id}/settings/rum`

Operation ID: `web-analytics-toggle-rum`

Toggles RUM on/off for an existing zone.

## Definition

```yaml
{"operationId": "web-analytics-toggle-rum", "summary": "Toggle RUM on/off for a zone", "description": "Toggles RUM on/off for an existing zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_toggle-rum-request"}}}}, "responses": {"200": {"description": "Rum toggled on/off for an existing zone.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_rum-site-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.rum", "x-fern-sdk-method-name": "edit"}
```
