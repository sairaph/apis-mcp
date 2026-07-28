---
title: Fetch BGP Prefix
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-bgp-prefixes-bgp-p-6dc2fed8
path: operations/ip-address-management-bgp-prefixes
description: Retrieve a single BGP Prefix according to its identifier
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}
operation_ids:
    - ip-address-management-prefixes-fetch-bgp-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch BGP Prefix

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}`

Operation ID: `ip-address-management-prefixes-fetch-bgp-prefix`

Retrieve a single BGP Prefix according to its identifier

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-fetch-bgp-prefix", "summary": "Fetch BGP Prefix", "description": "Retrieve a single BGP Prefix according to its identifier", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "bgp_prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_bgp_prefix_identifier"}}], "responses": {"200": {"description": "Fetch BGP Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response_bgp"}}}}, "4XX": {"description": "Fetch BGP Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response_bgp"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management BGP Prefixes"], "x-api-token-group": ["Magic Transit Read", "Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: Read", "IP Prefixes: BGP On Demand Write", "IP Prefixes: BGP On Demand Read"]}
```
