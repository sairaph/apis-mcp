---
title: Patch Health Check
page_id: operation-patch-zones-zone-id-smart-shield-healthchecks-healthcheck-id-369e3635
path: operations/health-checks
description: Patch a configured health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}
operation_ids:
    - smart-shield-patch-health-check
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Health Check

`PATCH /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}`

Operation ID: `smart-shield-patch-health-check`

Patch a configured health check.

## Definition

```yaml
{"operationId": "smart-shield-patch-health-check", "summary": "Patch Health Check", "description": "Patch a configured health check.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_query_healthcheck"}}}}, "responses": {"200": {"description": "Patch Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_single_hc_response"}}}}, "4XX": {"description": "Patch Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_single_hc_response"}, {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
