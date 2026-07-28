---
title: Add a zone membership to an Address Map
page_id: operation-put-accounts-account-id-addressing-address-maps-address-map-id-zones-zon-756d0918
path: operations/ip-address-management-address-maps
description: Add a zone as a member of a particular address map.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}
operation_ids:
    - ip-address-management-address-maps-add-a-zone-membership-to-an-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a zone membership to an Address Map

`PUT /accounts/{account_id}/addressing/address_maps/{address_map_id}/zones/{zone_id}`

Operation ID: `ip-address-management-address-maps-add-a-zone-membership-to-an-address-map`

Add a zone as a member of a particular address map.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-add-a-zone-membership-to-an-address-map", "summary": "Add a zone membership to an Address Map", "description": "Add a zone as a member of a particular address map.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_zone_identifier"}}, {"name": "address_map_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_address_map_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Add a zone membership to an Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-collection"}}}}, "4XX": {"description": "Add a zone membership to an Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"]}
```
