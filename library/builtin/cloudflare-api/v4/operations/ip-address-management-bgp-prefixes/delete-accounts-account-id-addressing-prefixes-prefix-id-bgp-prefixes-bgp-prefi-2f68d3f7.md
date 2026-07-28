---
title: Delete BGP Prefix
page_id: operation-delete-accounts-account-id-addressing-prefixes-prefix-id-bgp-prefixes-bg-825f56ce
path: operations/ip-address-management-bgp-prefixes
description: Delete a BGP Prefix associated with the specified IP Prefix. A BGP Prefix must be withdrawn before it can be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}
operation_ids:
    - ip-address-management-prefixes-delete-bgp-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete BGP Prefix

`DELETE /accounts/{account_id}/addressing/prefixes/{prefix_id}/bgp/prefixes/{bgp_prefix_id}`

Operation ID: `ip-address-management-prefixes-delete-bgp-prefix`

Delete a BGP Prefix associated with the specified IP Prefix. A BGP Prefix must be withdrawn before it can be deleted.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-delete-bgp-prefix", "summary": "Delete BGP Prefix", "description": "Delete a BGP Prefix associated with the specified IP Prefix. A BGP Prefix must be withdrawn before it can be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "bgp_prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_bgp_prefix_identifier"}}], "responses": {"200": {"description": "Delete BGP Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common"}}}}, "4XX": {"description": "Delete BGP Prefix response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management BGP Prefixes"]}
```
