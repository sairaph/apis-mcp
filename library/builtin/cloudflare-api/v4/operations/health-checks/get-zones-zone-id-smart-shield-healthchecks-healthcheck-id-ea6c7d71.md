---
title: Health Check Details
page_id: operation-get-zones-zone-id-smart-shield-healthchecks-healthcheck-id-270f2c71
path: operations/health-checks
description: Fetch a single configured health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}
operation_ids:
    - smart-shield-health-check-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Health Check Details

`GET /zones/{zone_id}/smart_shield/healthchecks/{healthcheck_id}`

Operation ID: `smart-shield-health-check-details`

Fetch a single configured health check.

## Definition

```yaml
{"operationId": "smart-shield-health-check-details", "summary": "Health Check Details", "description": "Fetch a single configured health check.", "parameters": [{"name": "healthcheck_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/smartshield_identifier"}}], "responses": {"200": {"description": "Health Check Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/smartshield_single_hc_response"}}}}, "4XX": {"description": "Health Check Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/smartshield_single_hc_response"}, {"$ref": "#/components/schemas/smartshield_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Health Checks"], "x-api-token-group": ["Health Checks Write", "Health Checks Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
