---
title: Create BGP Prefix
page_id: operation-post-accounts-account-id-addressing-prefixes-prefix-id-bgp-prefixes-a16cb760
path: operations/ip-address-management-bgp-prefixes
description: Create a BGP prefix, controlling the BGP advertisement status of a specific subnet. When created, BGP prefixes are initially withdrawn, and can be advertised with the Update BGP Prefix API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes
operation_ids:
    - ip-address-management-prefixes-create-bgp-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create BGP Prefix

`POST /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes`

Operation ID: `ip-address-management-prefixes-create-bgp-prefix`

Create a BGP prefix, controlling the BGP advertisement status of a specific subnet. When created, BGP prefixes are initially withdrawn, and can be advertised with the Update BGP Prefix API.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-create-bgp-prefix", "summary": "Create BGP Prefix", "description": "Create a BGP prefix, controlling the BGP advertisement status of a specific subnet. When created, BGP prefixes are initially withdrawn, and can be advertised with the Update BGP Prefix API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_bgp_prefix_create"}}}}, "responses": {"200": {"description": "Create BGP Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response_bgp"}}}}, "4XX": {"description": "Create BGP Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response_bgp"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management BGP Prefixes"]}
```
