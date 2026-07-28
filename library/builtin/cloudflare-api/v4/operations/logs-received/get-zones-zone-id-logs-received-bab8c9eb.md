---
title: Get logs received
page_id: operation-get-zones-zone-id-logs-received-4baa0791
path: operations/logs-received
description: 'The `/received` api route allows customers to retrieve their edge HTTP logs. The basic access pattern is "give me all the logs for zone Z for minute M", where the minute M refers to the time records were received at Cloudflare''s central data center. `start` is inclusive, and `end` is exclusive. Because of that, to get all data, at minutely cadence, starting at 10AM, the proper values are: `start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z`, then `start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z` and so on; the overlap will be handled properly.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/received
operation_ids:
    - get-zones-zone_id-logs-received
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get logs received

`GET /zones/{zone_id}/logs/received`

Operation ID: `get-zones-zone_id-logs-received`

The `/received` api route allows customers to retrieve their edge HTTP logs. The basic access pattern is "give me all the logs for zone Z for minute M", where the minute M refers to the time records were received at Cloudflare's central data center. `start` is inclusive, and `end` is exclusive. Because of that, to get all data, at minutely cadence, starting at 10AM, the proper values are: `start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z`, then `start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z` and so on; the overlap will be handled properly.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logs-received", "summary": "Get logs received", "description": "The `/received` api route allows customers to retrieve their edge HTTP logs. The basic access pattern is \"give me all the logs for zone Z for minute M\", where the minute M refers to the time records were received at Cloudflare's central data center. `start` is inclusive, and `end` is exclusive. Because of that, to get all data, at minutely cadence, starting at 10AM, the proper values are: `start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z`, then `start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z` and so on; the overlap will be handled properly.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logshare_identifier"}}, {"name": "start", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_start"}}, {"name": "end", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/logshare_end"}}, {"name": "fields", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_fields"}}, {"name": "sample", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_sample"}}, {"name": "count", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_count"}}, {"name": "timestamps", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_timestamps"}}], "responses": {"200": {"description": "Get logs received response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logshare_logs_response_json_lines"}}}}, "4XX": {"description": "Get logs received response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logshare_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logs Received"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-cfPermissionsRequired": {"enum": ["#logs:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
