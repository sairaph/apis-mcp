---
title: Start page test
page_id: operation-post-zones-zone-id-speed-api-pages-url-tests-93b84276
path: operations/observatory
description: Starts a test for a specific webpage, in a specific region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/speed_api/pages/{url}/tests
operation_ids:
    - speed-create-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start page test

`POST /zones/{zone_id}/speed_api/pages/{url}/tests`

Operation ID: `speed-create-test`

Starts a test for a specific webpage, in a specific region.

## Definition

```yaml
{"operationId": "speed-create-test", "summary": "Start page test", "description": "Starts a test for a specific webpage, in a specific region.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"region": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}}}}}}, "responses": {"200": {"description": "Page test details.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_page-test-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages.tests", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
