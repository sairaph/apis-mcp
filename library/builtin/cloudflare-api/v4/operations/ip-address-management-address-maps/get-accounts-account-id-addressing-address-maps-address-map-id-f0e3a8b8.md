---
title: Address Map Details
page_id: operation-get-accounts-account-id-addressing-address-maps-address-map-id-f734fe74
path: operations/ip-address-management-address-maps
description: Show a particular address map owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}
operation_ids:
    - ip-address-management-address-maps-address-map-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Address Map Details

`GET /accounts/{account_id}/addressing/address_maps/{address_map_id}`

Operation ID: `ip-address-management-address-maps-address-map-details`

Show a particular address map owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-address-map-details", "summary": "Address Map Details", "description": "Show a particular address map owned by the account.", "parameters": [{"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "Address Map Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_full_response"}}}}, "4XX": {"description": "Address Map Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_full_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write", "Address Maps Read"]}
```
