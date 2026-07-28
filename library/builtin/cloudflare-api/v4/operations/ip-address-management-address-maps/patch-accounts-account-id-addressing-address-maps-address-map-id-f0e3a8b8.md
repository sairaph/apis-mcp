---
title: Update Address Map
page_id: operation-patch-accounts-account-id-addressing-address-maps-address-map-id-5976f670
path: operations/ip-address-management-address-maps
description: Modify properties of an address map owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}
operation_ids:
    - ip-address-management-address-maps-update-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Address Map

`PATCH /accounts/{account_id}/addressing/address_maps/{address_map_id}`

Operation ID: `ip-address-management-address-maps-update-address-map`

Modify properties of an address map owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-update-address-map", "summary": "Update Address Map", "description": "Modify properties of an address map owned by the account.", "parameters": [{"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"default_sni": {"$ref": "#/components/schemas/addressing_default_sni"}, "description": {"$ref": "#/components/schemas/addressing_schemas-description"}, "enabled": {"$ref": "#/components/schemas/addressing_enabled"}}}}}}, "responses": {"200": {"description": "Update Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_components-schemas-single_response"}}}}, "4XX": {"description": "Update Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_components-schemas-single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write"]}
```
