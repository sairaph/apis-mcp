---
title: Get log retention flag
page_id: operation-get-zones-zone-id-logs-control-retention-flag-799be06d
path: operations/logs-received
description: Gets log retention flag for Logpull API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/control/retention/flag
operation_ids:
    - get-zones-zone_id-logs-control-retention-flag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get log retention flag

`GET /zones/{zone_id}/logs/control/retention/flag`

Operation ID: `get-zones-zone_id-logs-control-retention-flag`

Gets log retention flag for Logpull API.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logs-control-retention-flag", "summary": "Get log retention flag", "description": "Gets log retention flag for Logpull API.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logcontrol_identifier"}}], "responses": {"200": {"description": "Get log retention flag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logcontrol_retention_flag_response_single"}}}}, "4XX": {"description": "Get log retention flag response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logcontrol_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logs Received"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read", "#analytics:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
