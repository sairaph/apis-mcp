---
title: Update Health Check
page_id: operation-put-zones-zone-id-healthchecks-healthcheck-id-6de1e1fa
path: operations/health-checks
description: Update a configured health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/healthchecks/{healthcheck_id}
operation_ids:
    - health-checks-update-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Health Check

`PUT /zones/{zone_id}/healthchecks/{healthcheck_id}`

Operation ID: `health-checks-update-health-check`

Update a configured health check.

## Definition

```yaml
{"operationId": "health-checks-update-health-check", "summary": "Update Health Check", "description": "Update a configured health check.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_query_healthcheck"}}}}, "responses": {"200": {"description": "Update Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_single_response"}}}}, "4XX": {"description": "Update Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/healthchecks_single_response"}, {"$ref": "#/components/schemas/healthchecks_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
