---
title: Create scheduled page test
page_id: operation-post-zones-zone-id-speed-api-schedule-url-c3b8bdeb
path: operations/observatory
description: Creates a scheduled test for a page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/speed_api/schedule/{url}
operation_ids:
    - speed-create-scheduled-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create scheduled page test

`POST /zones/{zone_id}/speed_api/schedule/{url}`

Operation ID: `speed-create-scheduled-test`

Creates a scheduled test for a page.

## Definition

```yaml
{"operationId": "speed-create-scheduled-test", "summary": "Create scheduled page test", "description": "Creates a scheduled test for a page.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "region", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}, "x-stainless-terraform-configurability": "computed_optional"}, {"name": "frequency", "in": "query", "description": "The frequency of the scheduled test. Defaults to WEEKLY for free plans, DAILY for paid plans.", "schema": {"$ref": "#/components/schemas/observatory_schedule_frequency"}, "x-stainless-terraform-configurability": "computed_optional"}], "responses": {"200": {"description": "Page test schedule.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_create-schedule-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.schedule", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
