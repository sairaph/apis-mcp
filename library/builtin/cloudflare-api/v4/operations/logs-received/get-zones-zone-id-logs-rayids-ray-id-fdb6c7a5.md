---
title: Get logs RayIDs
page_id: operation-get-zones-zone-id-logs-rayids-ray-id-070f63c2
path: operations/logs-received
description: The `/rayids` api route allows lookups by specific rayid. The rayids route will return zero, one, or more records (ray ids are not unique).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/rayids/{ray_id}
operation_ids:
    - get-zones-zone_id-logs-rayids-ray_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get logs RayIDs

`GET /zones/{zone_id}/logs/rayids/{ray_id}`

Operation ID: `get-zones-zone_id-logs-rayids-ray_id`

The `/rayids` api route allows lookups by specific rayid. The rayids route will return zero, one, or more records (ray ids are not unique).

## Definition

```yaml
{"operationId": "get-zones-zone_id-logs-rayids-ray_id", "summary": "Get logs RayIDs", "description": "The `/rayids` api route allows lookups by specific rayid. The rayids route will return zero, one, or more records (ray ids are not unique).", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logshare_identifier"}}, {"name": "ray_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logshare_ray_identifier"}}, {"name": "fields", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_fields"}}, {"name": "timestamps", "in": "query", "schema": {"$ref": "#/components/schemas/logshare_timestamps"}}], "responses": {"200": {"description": "Get logs RayIDs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logshare_logs_response_json_lines"}}}}, "4XX": {"description": "Get logs RayIDs response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logshare_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logs Received"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
