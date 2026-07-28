---
title: List BGP Prefixes
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-bgp-prefixes-6bdd29df
path: operations/ip-address-management-bgp-prefixes
description: List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to control which specific subnets are advertised to the Internet. It is possible to advertise subnets more specific than an IP Prefix by creating more specific BGP Prefixes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes
operation_ids:
    - ip-address-management-prefixes-list-bgp-prefixes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List BGP Prefixes

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes`

Operation ID: `ip-address-management-prefixes-list-bgp-prefixes`

List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to control which specific subnets are advertised to the Internet. It is possible to advertise subnets more specific than an IP Prefix by creating more specific BGP Prefixes.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-list-bgp-prefixes", "summary": "List BGP Prefixes", "description": "List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to control which specific subnets are advertised to the Internet. It is possible to advertise subnets more specific than an IP Prefix by creating more specific BGP Prefixes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}], "responses": {"200": {"description": "List BGP Prefixes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_response_collection_bgp"}}}}, "4XX": {"description": "List BGP Prefixes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_response_collection_bgp"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management BGP Prefixes"], "x-api-token-group": ["Magic Transit Read", "Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: Read", "IP Prefixes: BGP On Demand Write", "IP Prefixes: BGP On Demand Read"]}
```
