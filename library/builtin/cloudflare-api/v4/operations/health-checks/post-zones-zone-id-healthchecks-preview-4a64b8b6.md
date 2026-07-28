---
title: Create Preview Health Check
page_id: operation-post-zones-zone-id-healthchecks-preview-89602f87
path: operations/health-checks
description: Create a new preview health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/healthchecks/preview
operation_ids:
    - health-checks-create-preview-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Preview Health Check

`POST /zones/{zone_id}/healthchecks/preview`

Operation ID: `health-checks-create-preview-health-check`

Create a new preview health check.

## Definition

```yaml
{"operationId": "health-checks-create-preview-health-check", "summary": "Create Preview Health Check", "description": "Create a new preview health check.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_query_healthcheck"}}}}, "responses": {"200": {"description": "Create Preview Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_single_response"}}}}, "4XX": {"description": "Create Preview Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/healthchecks_single_response"}, {"$ref": "#/components/schemas/healthchecks_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
