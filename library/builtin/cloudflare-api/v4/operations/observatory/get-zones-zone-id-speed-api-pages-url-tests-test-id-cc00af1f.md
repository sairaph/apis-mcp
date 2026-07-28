---
title: Get a page test result
page_id: operation-get-zones-zone-id-speed-api-pages-url-tests-test-id-7c6861c5
path: operations/observatory
description: Retrieves the result of a specific test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/pages/{url}/tests/{test_id}
operation_ids:
    - speed-get-test
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a page test result

`GET /zones/{zone_id}/speed_api/pages/{url}/tests/{test_id}`

Operation ID: `speed-get-test`

Retrieves the result of a specific test.

## Definition

```yaml
{"operationId": "speed-get-test", "summary": "Get a page test result", "description": "Retrieves the result of a specific test.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "test_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Page test result.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_page-test-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages.tests", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
