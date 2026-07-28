---
title: Delete Zone
page_id: operation-delete-zones-zone-id-0da0a0b8
path: operations/zone
description: Deletes an existing zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}
operation_ids:
    - zones-0-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Zone

`DELETE /zones/{zone_id}`

Operation ID: `zones-0-delete`

Deletes an existing zone.

## Definition

```yaml
{"operationId": "zones-0-delete", "summary": "Delete Zone", "description": "Deletes an existing zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}], "requestBody": {"content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Zone response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-single-id"}}}}, "4XX": {"description": "Delete Zone response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones", "x-fern-sdk-method-name": "delete"}
```
