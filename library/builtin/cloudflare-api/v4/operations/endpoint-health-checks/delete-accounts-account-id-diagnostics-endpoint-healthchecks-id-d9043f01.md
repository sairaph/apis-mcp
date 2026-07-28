---
title: Delete Endpoint Health Check
page_id: operation-delete-accounts-account-id-diagnostics-endpoint-healthchecks-id-7cb2b801
path: operations/endpoint-health-checks
description: Delete Endpoint Health Check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}
operation_ids:
    - diagnostics-endpoint-healthcheck-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Endpoint Health Check

`DELETE /accounts/{account_id}/diagnostics/endpoint-healthchecks/{id}`

Operation ID: `diagnostics-endpoint-healthcheck-delete`

Delete Endpoint Health Check.

## Definition

```yaml
{"operationId": "diagnostics-endpoint-healthcheck-delete", "summary": "Delete Endpoint Health Check", "description": "Delete Endpoint Health Check.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_identifier"}}, {"name": "id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-transit_uuid"}}], "responses": {"200": {"description": "Endpoint Health Checks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-transit_api-response-common"}}}}, "4XX": {"description": "Endpoint Health Check failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-transit_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Endpoint Health Checks"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
