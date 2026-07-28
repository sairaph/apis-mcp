---
title: Delete all page tests
page_id: operation-delete-zones-zone-id-speed-api-pages-url-tests-c5df70f1
path: operations/observatory
description: Deletes all tests for a specific webpage from a specific region. Deleted tests are still counted as part of the quota.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/speed_api/pages/{url}/tests
operation_ids:
    - speed-delete-tests
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all page tests

`DELETE /zones/{zone_id}/speed_api/pages/{url}/tests`

Operation ID: `speed-delete-tests`

Deletes all tests for a specific webpage from a specific region. Deleted tests are still counted as part of the quota.

## Definition

```yaml
{"operationId": "speed-delete-tests", "summary": "Delete all page tests", "description": "Deletes all tests for a specific webpage from a specific region. Deleted tests are still counted as part of the quota.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}, {"name": "url", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_url"}}, {"name": "region", "in": "query", "schema": {"allOf": [{"$ref": "#/components/schemas/observatory_region"}, {"default": "us-central1", "type": "string"}]}}], "responses": {"200": {"description": "Number of deleted tests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_count-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages.tests", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
