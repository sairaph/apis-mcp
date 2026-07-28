---
title: Delete Preview Health Check
page_id: operation-delete-zones-zone-id-healthchecks-preview-healthcheck-id-3a0b23ff
path: operations/health-checks
description: Delete a health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/healthchecks/preview/{healthcheck_id}
operation_ids:
    - health-checks-delete-preview-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Preview Health Check

`DELETE /zones/{zone_id}/healthchecks/preview/{healthcheck_id}`

Operation ID: `health-checks-delete-preview-health-check`

Delete a health check.

## Definition

```yaml
{"operationId": "health-checks-delete-preview-health-check", "summary": "Delete Preview Health Check", "description": "Delete a health check.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/healthchecks_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Preview Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/healthchecks_id_response"}}}}, "4XX": {"description": "Delete Preview Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/healthchecks_id_response"}, {"$ref": "#/components/schemas/healthchecks_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
