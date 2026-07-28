---
title: Update BGP Prefix
page_id: operation-patch-accounts-account-id-addressing-prefixes-prefix-id-bgp-prefixes-bgp-4e6500f0
path: operations/ip-address-management-bgp-prefixes
description: Update the properties of a BGP Prefix, such as the on demand advertisement status (advertised or withdrawn).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}
operation_ids:
    - ip-address-management-prefixes-update-bgp-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update BGP Prefix

`PATCH /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}`

Operation ID: `ip-address-management-prefixes-update-bgp-prefix`

Update the properties of a BGP Prefix, such as the on demand advertisement status (advertised or withdrawn).

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-update-bgp-prefix", "summary": "Update BGP Prefix", "description": "Update the properties of a BGP Prefix, such as the on demand advertisement status (advertised or withdrawn).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "bgp_prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_bgp_prefix_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_bgp_prefix_update_advertisement"}}}}, "responses": {"200": {"description": "Update BGP Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response_bgp"}}}}, "4XX": {"description": "Update BGP Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response_bgp"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management BGP Prefixes"], "x-api-token-group": ["Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: BGP On Demand Write"]}
```
