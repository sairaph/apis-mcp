---
title: Create Address Map
page_id: operation-post-accounts-account-id-addressing-address-maps-ee3a334a
path: operations/ip-address-management-address-maps
description: Create a new address map under the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/address_maps
operation_ids:
    - ip-address-management-address-maps-create-address-map
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Address Map

`POST /accounts/{account_id}/addressing/address_maps`

Operation ID: `ip-address-management-address-maps-create-address-map`

Create a new address map under the account.

## Definition

```yaml
{"operationId": "ip-address-management-address-maps-create-address-map", "summary": "Create Address Map", "description": "Create a new address map under the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"description": {"$ref": "#/components/schemas/addressing_schemas-description"}, "enabled": {"$ref": "#/components/schemas/addressing_enabled"}, "ips": {"type": "array", "items": {"$ref": "#/components/schemas/addressing_ip_address"}, "example": ["192.0.2.1"]}, "memberships": {"$ref": "#/components/schemas/addressing_membership_requests"}}}}}}, "responses": {"200": {"description": "Create Address Map response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_full_response"}}}}, "4XX": {"description": "Create Address Map response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_full_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Address Maps"], "x-api-token-group": ["Address Maps Write"]}
```
