---
title: Get Endpoint Health Check
page_id: operation-get-accounts-account-id-diagnostics-endpoint-healthchecks-id-5b7671db
path: operations/endpoint-health-checks
description: Get a single Endpoint Health Check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
operation_ids:
    - diagnostics-endpoint-healthcheck-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Endpoint Health Check

`GET /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}`

Operation ID: `diagnostics-endpoint-healthcheck-get`

Get a single Endpoint Health Check.

## Definition

```yaml
{"operationId": "diagnostics-endpoint-healthcheck-get", "summary": "Get Endpoint Health Check", "description": "Get a single Endpoint Health Check.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}, {"name": "id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_uuid"}}], "responses": {"200": {"description": "Endpoint Health Checks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check_response_single"}}}}, "4XX": {"description": "Endpoint Health Check failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_traceroute_response_collection"}, {"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Endpoint Health Checks"], "x-api-token-group": ["Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
