---
title: Create Health Check
page_id: operation-post-zones-zone-id-smart-shield-healthchecks-3cdd188d
path: operations/health-checks
description: Create a new health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/smart_shield/healthchecks
operation_ids:
    - smart-shield-create-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Health Check

`POST /zones/{zone_id}/smart_shield/healthchecks`

Operation ID: `smart-shield-create-health-check`

Create a new health check.

## Definition

```yaml
{"operationId": "smart-shield-create-health-check", "summary": "Create Health Check", "description": "Create a new health check.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_query_healthcheck"}}}}, "responses": {"200": {"description": "Create Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_single_hc_response"}}}}, "4XX": {"description": "Create Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_single_hc_response"}, {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
