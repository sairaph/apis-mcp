---
title: List Address Maps
page_id: operation-get-accounts-account-id-addressing-address-maps-c01cb486
path: operations/ip-address-management-address-maps
description: List all address maps owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps
operation_ids:
    - ip-address-management-address-maps-list-address-maps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Address Maps

`GET /accounts/{account_id}/addressing/address_maps`

Operation ID: `ip-address-management-address-maps-list-address-maps`

List all address maps owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-list-address-maps", "summary": "List Address Maps", "description": "List all address maps owned by the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "List Address Maps response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_components-schemas-response_collection"}}}}, "4XX": {"description": "List Address Maps response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_components-schemas-response_collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write", "Address Maps Read"]}
```
