---
title: Endpoint Health Check
page_id: operation-post-accounts-account-id-diagnostics-endpoint-healthchecks-8c28ef24
path: operations/endpoint-health-checks
description: Create Endpoint Health Check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/diagnostics/endpoint-healthchecks
operation_ids:
    - diagnostics-endpoint-healthcheck-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Endpoint Health Check

`POST /accounts/{account_id}/diagnostics/endpoint-healthchecks`

Operation ID: `diagnostics-endpoint-healthcheck-create`

Create Endpoint Health Check.

## Definition

```yaml
{"operationId": "diagnostics-endpoint-healthcheck-create", "summary": "Endpoint Health Check", "description": "Create Endpoint Health Check.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check"}}}}, "responses": {"201": {"description": "Endpoint Health Check response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check_response_single"}}}}, "4XX": {"description": "Endpoint Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Endpoint Health Checks"], "x-api-token-group": ["Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
