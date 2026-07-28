---
title: List fields
page_id: operation-get-zones-zone-id-logs-received-fields-417169bd
path: operations/logs-received
description: Lists all fields available. The response is json object with key-value pairs, where keys are field names, and values are descriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/received/fields
operation_ids:
    - get-zones-zone_id-logs-received-fields
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List fields

`GET /zones/{zone_id}/logs/received/fields`

Operation ID: `get-zones-zone_id-logs-received-fields`

Lists all fields available. The response is json object with key-value pairs, where keys are field names, and values are descriptions.

## Definition

```yaml
{"operationId": "get-zones-zone_id-logs-received-fields", "summary": "List fields", "description": "Lists all fields available. The response is json object with key-value pairs, where keys are field names, and values are descriptions.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logshare_identifier"}}], "responses": {"200": {"description": "List fields response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logshare_fields_response"}}}}, "4XX": {"description": "List fields response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logshare_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logs Received"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
