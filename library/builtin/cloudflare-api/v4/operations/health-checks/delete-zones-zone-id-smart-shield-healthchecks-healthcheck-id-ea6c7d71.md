---
title: Delete Health Check
page_id: operation-delete-zones-zone-id-smart-shield-healthchecks-healthcheck-id-0d1a10fd
path: operations/health-checks
description: Delete a health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}
operation_ids:
    - smart-shield-delete-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Health Check

`DELETE /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}`

Operation ID: `smart-shield-delete-health-check`

Delete a health check.

## Definition

```yaml
{"operationId": "smart-shield-delete-health-check", "summary": "Delete Health Check", "description": "Delete a health check.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_single_hc_id_response"}}}}, "4XX": {"description": "Delete Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_single_hc_id_response"}, {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
