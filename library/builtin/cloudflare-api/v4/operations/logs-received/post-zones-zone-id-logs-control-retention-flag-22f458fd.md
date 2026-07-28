---
title: Update log retention flag
page_id: operation-post-zones-zone-id-logs-control-retention-flag-47854a1d
path: operations/logs-received
description: Updates log retention flag for Logpull API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logs/control/retention/flag
operation_ids:
    - post-zones-zone_id-logs-control-retention-flag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update log retention flag

`POST /zones/{zone_id}/logs/control/retention/flag`

Operation ID: `post-zones-zone_id-logs-control-retention-flag`

Updates log retention flag for Logpull API.

## Definition

```yaml
{"operationId": "post-zones-zone_id-logs-control-retention-flag", "summary": "Update log retention flag", "description": "Updates log retention flag for Logpull API.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logcontrol_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_retention_flag"}}}}, "responses": {"200": {"description": "Update log retention flag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_retention_flag_response_single"}}}}, "4XX": {"description": "Update log retention flag response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logs Received"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
