---
title: List Prefixes
page_id: operation-get-accounts-account-id-addressing-prefixes-6ba6d361
path: operations/ip-address-management-prefixes
description: List all prefixes owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes
operation_ids:
    - ip-address-management-prefixes-list-prefixes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Prefixes

`GET /accounts/{account_id}/addressing/prefixes`

Operation ID: `ip-address-management-prefixes-list-prefixes`

List all prefixes owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-list-prefixes", "summary": "List Prefixes", "description": "List all prefixes owned by the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "List Prefixes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_response_collection"}}}}, "4XX": {"description": "List Prefixes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_response_collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Read", "Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: Read", "IP Prefixes: BGP On Demand Write", "IP Prefixes: BGP On Demand Read"]}
```
