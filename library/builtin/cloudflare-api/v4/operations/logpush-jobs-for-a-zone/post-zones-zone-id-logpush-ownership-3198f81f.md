---
title: Get ownership challenge
page_id: operation-post-zones-zone-id-logpush-ownership-06831f98
path: operations/logpush-jobs-for-a-zone
description: Gets a new ownership challenge sent to your destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logpush/ownership
operation_ids:
    - post-zones-zone_id-logpush-ownership
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ownership challenge

`POST /zones/{zone_id}/logpush/ownership`

Operation ID: `post-zones-zone_id-logpush-ownership`

Gets a new ownership challenge sent to your destination.

## Definition

```yaml
{"operationId": "post-zones-zone_id-logpush-ownership", "summary": "Get ownership challenge", "description": "Gets a new ownership challenge sent to your destination.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination_conf": {"$ref": "#/components/schemas/logpush_destination_conf"}}, "required": ["destination_conf"]}}}}, "responses": {"200": {"description": "Get ownership challenge response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_get_ownership_response"}}}}, "4XX": {"description": "Get ownership challenge response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for a zone"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.ownership", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
