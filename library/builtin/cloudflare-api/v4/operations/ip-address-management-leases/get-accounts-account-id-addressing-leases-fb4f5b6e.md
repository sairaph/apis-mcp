---
title: List Leases
page_id: operation-get-accounts-account-id-addressing-leases-baa3c0d3
path: operations/ip-address-management-leases
description: List all leases owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/leases
operation_ids:
    - ip-address-management-list-leases
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Leases

`GET /accounts/{account_id}/addressing/leases`

Operation ID: `ip-address-management-list-leases`

List all leases owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-list-leases", "summary": "List Leases", "description": "List all leases owned by the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_schemas-account_identifier"}}], "responses": {"200": {"description": "List Leases response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_leases_components-schemas-response_collection"}}}}, "4XX": {"description": "List Leases response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_leases_components-schemas-response_collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Leases"], "x-api-token-group": null}
```
