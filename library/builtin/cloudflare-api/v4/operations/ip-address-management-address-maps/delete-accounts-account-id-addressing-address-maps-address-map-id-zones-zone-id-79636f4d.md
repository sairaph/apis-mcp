---
title: Remove a zone membership from an Address Map
page_id: operation-delete-accounts-account-id-addressing-address-maps-address-map-id-zones-7a83dc46
path: operations/ip-address-management-address-maps
description: Remove a zone as a member of a particular address map.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}
operation_ids:
    - ip-address-management-address-maps-remove-a-zone-membership-from-an-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove a zone membership from an Address Map

`DELETE /accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}`

Operation ID: `ip-address-management-address-maps-remove-a-zone-membership-from-an-address-map`

Remove a zone as a member of a particular address map.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-remove-a-zone-membership-from-an-address-map", "summary": "Remove a zone membership from an Address Map", "description": "Remove a zone as a member of a particular address map.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_zone_identifier"}}, {"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Remove a zone membership from an Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-collection"}}}}, "4XX": {"description": "Remove a zone membership from an Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write"]}
```
