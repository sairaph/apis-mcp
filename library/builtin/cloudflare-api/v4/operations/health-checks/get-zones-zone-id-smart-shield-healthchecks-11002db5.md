---
title: List Health Checks
page_id: operation-get-zones-zone-id-smart-shield-healthchecks-67ae34f0
path: operations/health-checks
description: List configured health checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/smart_shield/healthchecks
operation_ids:
    - smart-shield-list-health-checks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Health Checks

`GET /zones/{zone_id}/smart_shield/healthchecks`

Operation ID: `smart-shield-list-health-checks`

List configured health checks.

## Definition

```yaml
{"operationId": "smart-shield-list-health-checks", "summary": "List Health Checks", "description": "List configured health checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}, {"$ref": "#/components/parameters/smartshield_page"}, {"$ref": "#/components/parameters/smartshield_per_page"}], "responses": {"200": {"description": "List Health Checks response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_response_collection"}}}}, "4XX": {"description": "List Health Checks response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_response_collection"}, {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write", "Health Checks Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
