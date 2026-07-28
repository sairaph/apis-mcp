---
title: Health Check Preview Details
page_id: operation-get-zones-zone-id-healthchecks-preview-healthcheck-id-e7a144ad
path: operations/health-checks
description: Fetch a single configured health check preview.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/healthchecks/preview/{healthcheck_id}
operation_ids:
    - health-checks-health-check-preview-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Health Check Preview Details

`GET /zones/{zone_id}/healthchecks/preview/{healthcheck_id}`

Operation ID: `health-checks-health-check-preview-details`

Fetch a single configured health check preview.

## Definition

```yaml
{"operationId": "health-checks-health-check-preview-details", "summary": "Health Check Preview Details", "description": "Fetch a single configured health check preview.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}], "responses": {"200": {"description": "Health Check Preview Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_single_response"}}}}, "4XX": {"description": "Health Check Preview Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/healthchecks_single_response"}, {"$ref": "#/components/schemas/healthchecks_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write", "Health Checks Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
