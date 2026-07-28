---
title: Delete scheduled page test
page_id: operation-delete-zones-zone-id-speed-api-schedule-url-0f595b93
path: operations/observatory
description: Deletes a scheduled test for a page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/speed_api/schedule/{url}
operation_ids:
    - speed-delete-test-schedule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete scheduled page test

`DELETE /zones/{zone_id}/speed_api/schedule/{url}`

Operation ID: `speed-delete-test-schedule`

Deletes a scheduled test for a page.

## Definition

```yaml
{"operationId": "speed-delete-test-schedule", "summary": "Delete scheduled page test", "description": "Deletes a scheduled test for a page.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "region", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}}], "responses": {"200": {"description": "Number of deleted tests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_count-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.schedule", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
