---
title: Remove an IP from an Address Map
page_id: operation-delete-accounts-account-id-addressing-address-maps-address-map-id-ips-ip-0ebe7d78
path: operations/ip-address-management-address-maps
description: Remove an IP from a particular address map.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}/ips/{ip_address}
operation_ids:
    - ip-address-management-address-maps-remove-an-ip-from-an-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove an IP from an Address Map

`DELETE /accounts/{account_id}/addressing/address_maps/{address_map_id}/ips/{ip_address}`

Operation ID: `ip-address-management-address-maps-remove-an-ip-from-an-address-map`

Remove an IP from a particular address map.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-remove-an-ip-from-an-address-map", "summary": "Remove an IP from an Address Map", "description": "Remove an IP from a particular address map.", "parameters": [{"name": "ip_address", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_ip_address"}}, {"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Remove an IP from an Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-collection"}}}}, "4XX": {"description": "Remove an IP from an Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write"]}
```
