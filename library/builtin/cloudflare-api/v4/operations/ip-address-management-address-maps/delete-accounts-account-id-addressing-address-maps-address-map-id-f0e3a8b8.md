---
title: Delete Address Map
page_id: operation-delete-accounts-account-id-addressing-address-maps-address-map-id-2a248760
path: operations/ip-address-management-address-maps
description: Delete a particular address map owned by the account. An Address Map must be disabled before it can be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}
operation_ids:
    - ip-address-management-address-maps-delete-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Address Map

`DELETE /accounts/{account_id}/addressing/address_maps/{address_map_id}`

Operation ID: `ip-address-management-address-maps-delete-address-map`

Delete a particular address map owned by the account. An Address Map must be disabled before it can be deleted.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-delete-address-map", "summary": "Delete Address Map", "description": "Delete a particular address map owned by the account. An Address Map must be disabled before it can be deleted.", "parameters": [{"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-collection"}}}}, "4XX": {"description": "Delete Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write"]}
```
