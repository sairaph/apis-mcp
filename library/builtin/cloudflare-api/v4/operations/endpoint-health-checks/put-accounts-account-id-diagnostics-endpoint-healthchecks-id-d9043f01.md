---
title: Update Endpoint Health Check
page_id: operation-put-accounts-account-id-diagnostics-endpoint-healthchecks-id-e32638a5
path: operations/endpoint-health-checks
description: Update a Endpoint Health Check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
operation_ids:
    - diagnostics-endpoint-healthcheck-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Endpoint Health Check

`PUT /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}`

Operation ID: `diagnostics-endpoint-healthcheck-update`

Update a Endpoint Health Check.

## Definition

```yaml
{"operationId": "diagnostics-endpoint-healthcheck-update", "summary": "Update Endpoint Health Check", "description": "Update a Endpoint Health Check.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}, {"name": "id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check"}}}}, "responses": {"200": {"description": "Endpoint Health Checks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_endpoint_health_check_response_single"}}}}, "4XX": {"description": "Endpoint Health Check failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Endpoint Health Checks"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
