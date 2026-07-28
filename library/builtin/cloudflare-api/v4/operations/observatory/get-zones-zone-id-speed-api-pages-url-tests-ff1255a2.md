---
title: List page test history
page_id: operation-get-zones-zone-id-speed-api-pages-url-tests-79d5674b
path: operations/observatory
description: Test history (list of tests) for a specific webpage.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/pages/{url}/tests
operation_ids:
    - speed-list-test-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List page test history

`GET /zones/{zone_id}/speed_api/pages/{url}/tests`

Operation ID: `speed-list-test-history`

Test history (list of tests) for a specific webpage.

## Definition

```yaml
{"operationId": "speed-list-test-history", "summary": "List page test history", "description": "Test history (list of tests) for a specific webpage.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "region", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}}], "responses": {"200": {"description": "List of test history for a page.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_page-test-response-collection"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages.tests", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
