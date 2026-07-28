---
title: List Health Checks
page_id: operation-get-zones-zone-id-healthchecks-b4dcd543
path: operations/health-checks
description: List configured health checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/healthchecks
operation_ids:
    - health-checks-list-health-checks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Health Checks

`GET /zones/{zone_id}/healthchecks`

Operation ID: `health-checks-list-health-checks`

List configured health checks.

## Definition

```yaml
{"operationId": "health-checks-list-health-checks", "summary": "List Health Checks", "description": "List configured health checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}, {"$ref": "#/components/parameters/healthchecks_page"}, {"$ref": "#/components/parameters/healthchecks_per_page"}], "responses": {"200": {"description": "List Health Checks response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_response_collection"}}}}, "4XX": {"description": "List Health Checks response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/healthchecks_response_collection"}, {"$ref": "#/components/schemas/healthchecks_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write", "Health Checks Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
