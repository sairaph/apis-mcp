---
title: List Endpoint Health Checks
page_id: operation-get-accounts-account-id-diagnostics-endpoint-healthchecks-64fd02b3
path: operations/endpoint-health-checks
description: List Endpoint Health Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/diagnostics/endpoint-healthchecks
operation_ids:
    - diagnostics-endpoint-healthcheck-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Endpoint Health Checks

`GET /accounts/{account_id}/diagnostics/endpoint-healthchecks`

Operation ID: `diagnostics-endpoint-healthcheck-list`

List Endpoint Health Checks.

## Definition

```yaml
{"operationId": "diagnostics-endpoint-healthcheck-list", "summary": "List Endpoint Health Checks", "description": "List Endpoint Health Checks.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}], "responses": {"200": {"description": "Endpoint Health Checks for account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check_response_single"}}}}, "4XX": {"description": "Endpoint Health Check response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Endpoint Health Checks"], "x-api-token-group": ["Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
